package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(something interface{}) {
	fmt.Println(GetJSON(something))
}

func GetJSON(something interface{}) string {
	if IsRunByAWS() {
		// log without intend
		enc, err := json.Marshal(something)
		if err != nil {
			panic(err)
		}
		return string(enc)
	}

	enc, err := json.MarshalIndent(something, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(enc)
}
