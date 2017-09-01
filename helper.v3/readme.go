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
> 仅含免费题

%s
## 参考解答
%s

%s
`
	headFormat := string(read("README_HEAD.md"))
	head := fmt.Sprintf(headFormat, lc.Username, lc.Username, lc.Ranking, lc.Username)

	acceptedPercent := lc.Categories[len(lc.Categories)-1].Total.Solved * 100 / lc.Categories[len(lc.Categories)-1].Total.Total

	count := lc.Categories.String()

	accepted := lc.Problems.acceptedString()

	tail := read("README_TAIL.md")
	content := fmt.Sprintf(template, head, acceptedPercent, count, accepted, tail)

	err := ioutil.WriteFile(file, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
