package main

import (
	"io/ioutil"
	"os"
)

func read(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}
