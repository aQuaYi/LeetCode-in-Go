package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func makeREADME(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	template := `%s

## 答题进度
%s

## 参考解答
%s

%s
`
	headFormat := string(read("README_HEAD.md"))
	head := fmt.Sprintf(headFormat, lc.Username, lc.Username, lc.Ranking, lc.Username)

	count := lc.Categories.String()

	solved := lc.Problems.String()

	tail := read("README_TAIL.md")
	content := fmt.Sprintf(template, head, count, solved, tail)

	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
