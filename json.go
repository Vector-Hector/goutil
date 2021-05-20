package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(something interface{}) {
	enc, err := json.MarshalIndent(something, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}
