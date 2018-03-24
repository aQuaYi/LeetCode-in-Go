package main

import (
	"io/ioutil"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// read 负责读取文件
// 这是一个通用的方法
func read(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}

// TODO: 编写一个写方法
