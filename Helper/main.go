package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 程序辅助设置
const (
	VERSION = "6.0.0"

	configFile      = "leetcode.toml"
	leetCodeFile    = "leetcode.json"
	unavailableFile = "unavailable.json"

	USAGE = `使用方法：
	1. helper [n] 	: 生成第 n 题的答题文件夹
	2. helper readme: 重新生成项目的 README.md 文件
`
)

// cfg 用于保存 LeetCode.com 的用户名和密码
// 程序中会有多处地方需要核对 用户名 ，所以设置成全局变量
var cfg config

func main() {
	fmt.Printf("helper version %s\n", VERSION)

	if len(os.Args) == 1 {
		fmt.Println(USAGE)
		return
	}

	// TODO: 在此生成 LeetCode 实例

	if os.Args[1] == "readme" {
		log.Println("~~ 开始重制 README.md 文档 ~~")
		rebuildReadme()
		log.Println("~~ 重制 README.md 完成 ~~")
		return
	}

	numProblem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("os.Args[1] 无法被转换成题号:", err)
	}

	log.Printf("~~ 开始生成第 %d 题的文件夹 ~~\n", numProblem)
	buildProblemDir(numProblem)
	log.Printf("~~ 第 %d 题的文件夹，已经生成 ~~\n", numProblem)
}
