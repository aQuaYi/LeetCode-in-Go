package main

import (
	"log"
	"os"
)

// 程序辅助设置
const (
	VERSION = "6.0.0"

	// unavailableFile = "unavailable.json"

	USAGE = `使用方法
	1. helper readme: 重新生成项目的 README.md 文件
	2. helper [n] 	: 生成第 n 题的答题文件夹
`
)

func main() {
	log.Printf("Hi, %s. I'm %s\n", getConfig().Username, VERSION)

	if len(os.Args) != 2 {
		log.Println(USAGE)
		return
	}

	switch os.Args[1] {
	case "readme":
		buildReadme()
	default:
		// buildProblemDir(os.Args[1])
	}

}
