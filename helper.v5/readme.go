package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func rebuildReadme() {
	// 由于网络原因，有时候 signin 比较慢
	signin()

	lc := lastestLeetCode()

	makeREADME(lc)
}

func makeREADME(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	// 更新 README.md 的内容
	template := `%s

## 答题进度：
%s
> 仅含能用 Go 语言解答的 Algorithms 类别的免费题
## 参考解答
%s
以下题目，暂时不能使用 Golang 解答
%s

%s
`

	head := getHead(lc)

	progressTable := lc.Algorithms.progressTable()

	acceptedTable := lc.AllProblems.accepted().table()

	unavailList := lc.UnavailProblem.list()

	tail := read("README_TAIL.md")

	content := fmt.Sprintf(template, head, acceptedPercent, count, accepted, tail)

	// 保存 README.md 文件
	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat, lc.Username, lc.Username, lc.Ranking, lc.Username)
}
