package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func HttpGet(url string, body interface{}) int {
	return request("GET", url, nil, body)
}

func HttpPost(url string, inputBody interface{}, outputBody interface{}) int {
	return request("POST", url, inputBody, outputBody)
}

func request(method string, url string, inputBody interface{}, outputBody interface{}) int {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	fmt.Println("->", method, url)

	var resp *http.Response
	switch method {
	case "GET":
		var err error
		resp, err = http.Get(url)
		if err != nil {
			if isHttpIgnorable(err) {
				return resp.StatusCode
			}
			panic(err)
		}
		break
	case "POST":
		inputBytes, err := json.Marshal(inputBody)
		if err != nil {panic(err)}
		resp, err = http.Post(url, "application/json", bytes.NewReader(inputBytes))
		if err != nil {
			if isHttpIgnorable(err) {
				return resp.StatusCode
			}
			panic(err)
		}
		break
	}
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http call resulted in", resp.StatusCode, "with body", string(buf.Bytes()))
	}
	err = json.Unmarshal(buf.Bytes(), outputBody)
	if err != nil {panic(err)}
	return resp.StatusCode
}

func isHttpIgnorable(err error) bool {
	errStr := err.Error()
	if strings.Contains(errStr, "EOF") {
		return true
	}
	return false
}

func DownloadFile(filepath string, url string) error {
	return DownloadFileWithHeaders(filepath, url, nil)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFileWithHeaders(filepath string, url string, headers map[string]string) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if headers != nil  {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	// Get the data
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	code := resp.StatusCode
	if code != 200 {
		return errors.New(fmt.Sprintf("%s %d", "request resulted in status code", code))
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
