package main

import (
	"fmt"
	"log"
	"os"
)

func buildReadme() {
	log.Println("开始，重建 README 文档")

	lc := newLeetCode()
	makeReadmeFile(lc)

	log.Println("完成，重建 README 文档")
}

func makeReadmeFile(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	// 更新 README.md 的内容
	template := `%s

## 进度

> 统计规则：1.免费题，2.算法题，3.能用 Go 解答

%s
## 题解

%s
以下免费的算法题，暂时不能使用 Go 解答

%s
#%s
`

	head := getHead(lc)

	progressTable := lc.Record.progressTable()

	availableTable := lc.Problems.available().table()

	unavailList := lc.Problems.unavailable().list()

	tail := read("README_TAIL.md")

	content := fmt.Sprintf(template, head, progressTable, availableTable, unavailList, tail)

	// 保存 README.md 文件
	write(file, content)
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat, lc.Username, lc.Ranking, lc.Username)
}
