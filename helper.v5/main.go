package main

import (
	"fmt"
	"log"
	"os"
)

// 程序辅助设置
const (
	VERSION = "5.7.1"

	configFile      = "leetcode.toml"
	leetCodeFile    = "leetcode.json"
	unavailableFile = "unavailable.json"
	tasksFile       = "tasks.txt"

	USAGE = `使用方法：
	1. helper readme : 重新生成项目的 README.md 文件
	2. helper n      : 生成第 n 题的答题文件夹 (暂时不可用)
	3. helper tasks  : 为未完成的题目，生成任务清单
	4. helper check  : 检查出不能用 Go 解答的题目，并保存题号到 unavailable.json
	5. helper v      : 查看 helper 的版本`
)

// cfg 用于保存 LeetCode.com 的用户名和密码
// 程序中会有多处地方需要核对 用户名 ，所以设置成全局变量
var cfg config

func main() {
	if len(os.Args) == 1 {
		fmt.Println(USAGE)
		fmt.Printf("helper version %s\n", VERSION)
		return
	}

	if os.Args[1] == "v" || os.Args[1] == "version" {
		fmt.Printf("helper version %s\n", VERSION)
		return
	}

	// 由于网络原因，有时候 signin 比较慢
	signin()

	switch os.Args[1] {
	case "h", "help":
		fmt.Println(USAGE)
	case "check":
		log.Println("~~ 开始检查剩余的题目 ~~")
		checkAvailable()
		log.Println("~~ 检查完成 ~~")
	case "readme", "r":
		log.Println("~~ 开始重制 README.md 文档 ~~")
		rebuildReadme()
		log.Println("~~ 重制 READ.md 完成 ~~")
	case "tasks", "t":
		log.Println("~~ 开始制作任务清单 ~~")
		printTasks()
		log.Println("~~ 制作任务清单完成 ~~")
	default:
		log.Println("~~ 开始生成答题文件夹 ~~")
		buildProblemDir(os.Args[1])
		log.Println("~~ 已经生成答题文件夹 ~~")
	}
}
