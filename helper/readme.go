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

## 答题进度 
%s

## 参考解答
%s

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
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return string(data)
}
