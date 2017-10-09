package main

import (
	"fmt"
	"os"
)

// 程序辅助设置
const (
	VERSION = "5.2.2"

	configFile      = "leetcode.toml"
	leetCodeFile    = "leetcode.json"
	unavailableFile = "unavailable.json"

	USAGE = `使用方法：
    1. helper readme : 重新生成项目的 README.md 文件。
    2. helper n      : 生成第 n 题的答题文件夹。
    3. helper -v     : 查看 helper 的版本`
)

// cfg 用于保存 LeetCode.com 的用户名和密码
// 程序中会有多处地方需要核对 用户名 ，所以设置成全局变量
var cfg config

func init() {
	// 由于网络原因，有时候 signin 比较慢
	signin()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(USAGE)
		return
	}

	switch os.Args[1] {
	case "-v", "-version":
		fmt.Printf("helper version %s\n", VERSION)
	case "-h", "-help":
		fmt.Println(USAGE)
	case "readme", "r":
		rebuildReadme()
	default:
		buildProblemDir(os.Args[1])
	}
}
