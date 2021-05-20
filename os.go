package util

import "os"

func Mkdir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
}
