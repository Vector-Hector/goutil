package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinPostToStruct(c *gin.Context, body interface{}) bool {
	bodyBuffer := new(bytes.Buffer)
	if _, err := bodyBuffer.ReadFrom(c.Request.Body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"data": `Couldn't read your input data.`,
		})
		return true
	}

	bodyBytes := bodyBuffer.Bytes()
	if err := json.Unmarshal(bodyBytes, body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"data": `Couldn't parse your input data.`,
		})
		return true
	}
	return false
}
