package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 程序辅助设置
const (
	VERSION = "3.0.1"
	USAGE   = `使用方法：
	1. 运行 helper 会重新生成项目的README.md。
	2. 运行 helper 123 会生成第123题的答题文件夹。`
)

var (
	problemNum = 0
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "-version":
			fmt.Printf("helper version %s\n", VERSION)
			return
		case "-h", "-help":
			fmt.Println(USAGE)
			return
		}

		var err error
		if problemNum, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalln(err)
		}
	}

	categories := []string{
		"Algorithms",
		// "database",
		"Draft",
		// "operating-system",
		// "shell",
		// "system-design",
	}

	update(categories)
}
