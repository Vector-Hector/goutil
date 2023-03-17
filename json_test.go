package util

import (
	"encoding/json"
	"testing"
)

func TestPrintJSON(t *testing.T) {
	str := `{
    "abc": "def"
}`
	var obj map[string]string
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		panic(err)
	}

	result := GetJSON(obj)
	if result != str {
		t.Error("json was not equal.")
	}
}
