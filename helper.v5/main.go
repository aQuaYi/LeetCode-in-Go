package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// 程序辅助设置
const (
	VERSION = "5.1.0"
	USAGE   = `使用方法：
    1. helper readme : 重新生成项目的 README.md 文件。
    2. helper n      : 生成第 n 题的答题文件夹。
    3. helper -v     : 查看 helper 的版本`
)

var cfg config

const (
	cfgFile         = "leetcode.toml"
	leetCodeFile    = "leetcode.json"
	unavailableFile = "unavailable.json"
)

func init() {
	// TODO: 这样真的有必要吗
	// 启动时，导入配置
	if _, err := toml.DecodeFile(cfgFile, &cfg); err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Hi, %s. \n", cfg.Login)
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
