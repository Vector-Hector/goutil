package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(something interface{}) {
	if IsRunByAWS() {
		// log without intend
		enc, err := json.Marshal(something)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(enc))
		return
	}
	
	enc, err := json.MarshalIndent(something, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}
