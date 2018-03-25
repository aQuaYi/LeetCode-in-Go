package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func buildReadme() {
	log.Println("~~ 开始重制 README.md 文档 ~~")

	lc := newLeetCode()
	lc.update()

	makeREADME(lc)

	log.Println("~~ 重制 README.md 完成 ~~")

	// 每更新一次，就保存一次
	lc.save()
}

func makeREADME(lc *leetcode) {
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
	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat, lc.Ranking, lc.Username)
}
