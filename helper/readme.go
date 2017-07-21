package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func makeREADME(c Categories, s Solveds) {
	file := "README.md"
	os.Remove(file)
	template := `%s

## 统计
%s

## 已完成的题目
%s

`
	head := readHead("README_HEAD.md")
	count := c.String()
	solved := s.String()
	content := fmt.Sprintf(template, head, count, solved)

	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func readHead(path string) string {
	head, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer head.Close()
	data, err := ioutil.ReadAll(head)
	return string(data)
}
