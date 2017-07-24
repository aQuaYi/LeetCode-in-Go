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
%s> 表中通过率是总体通过率。
%s
`
	head := read("README_HEAD.md")
	count := c.String()
	solved := s.String()
	tail := read("README_TAIL.md")
	content := fmt.Sprintf(template, head, count, solved, tail)

	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func read(path string) string {
	head, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer head.Close()

	data, err := ioutil.ReadAll(head)
	return string(data)
}
