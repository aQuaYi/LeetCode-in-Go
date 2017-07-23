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
>注意，上表中的通过率，是总体通过率，非本人的通过率。

## helper
[helper](./helper)会帮助处理大部分琐碎的工作。
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
