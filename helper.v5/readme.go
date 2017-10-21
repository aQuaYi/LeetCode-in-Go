package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func rebuildReadme() {
	lc := lastestLeetCode()
	makeREADME(lc)
}

func makeREADME(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	// 更新 README.md 的内容
	template := `%s

## 进度

%s
> 统计规则：1.免费题，2.算法题，3.能用 Go 解答

## 题解

%s

##以下免费的算法题，暂时不能使用 Go 解答

%s
#%s
`

	head := getHead(lc)

	progressTable := lc.Algorithms.progressTable()

	acceptedTable := lc.Problems.table()

	unavailList := lc.Problems.unavailable().list()

	tail := read("README_TAIL.md")

	content := fmt.Sprintf(template, head, progressTable, acceptedTable, unavailList, tail)

	// 保存 README.md 文件
	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat,
		lc.Username, lc.Username,
		lc.Progress, lc.Username,
		lc.Ranking, lc.Username,
	)
}
