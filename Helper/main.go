package main

import (
	"log"
	"os"
	"strconv"
)

// 程序辅助设置
const (
	VERSION = "6.0.0"

	// unavailableFile = "unavailable.json"

	USAGE = `使用方法：
	1. helper [n] 	: 生成第 n 题的答题文件夹
	2. helper readme: 重新生成项目的 README.md 文件
`
)

// cfg 用于保存 LeetCode.com 的用户名和密码
// 程序中会有多处地方需要核对 用户名 ，所以设置成全局变量
var cfg config

func main() {
	log.Printf("Hi, %s. I'm %s\n", getConfig().Username, VERSION)

	if len(os.Args) == 1 {
		log.Println(USAGE)
		return
	}

	if os.Args[1] == "readme" {
		rebuildReadme()
		return
	}

	numProblem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("os.Args[1] 无法被转换成题号:", err)
	}
	buildProblemDir(numProblem)
}
