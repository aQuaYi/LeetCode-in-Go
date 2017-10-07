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

## 答题进度：%d%%
统计规则
> 1. 免费题
> 2. 可以使用Go语言解答
%s
## 参考解答
%s

%s
`
	headFormat := string(read("README_HEAD.md"))
	head := fmt.Sprintf(headFormat, lc.Username, lc.Username, lc.Ranking, lc.Username)

	// 没有提供 Golang 解答方法的题目
	// TODO: 让这个方法自动化
	canNotSolve := 13
	acceptedPercent := lc.Categories[len(lc.Categories)-1].Total.Solved * 100 / (lc.Categories[len(lc.Categories)-1].Total.Total - canNotSolve)

	count := lc.Categories.String()

	accepted := lc.Problems.acceptedString()

	tail := read("README_TAIL.md")
	content := fmt.Sprintf(template, head, acceptedPercent, count, accepted, tail)

	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
