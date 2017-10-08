package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// 程序辅助设置
const (
	VERSION = "5.0.0"
	USAGE   = `使用方法：
	1. helper readme 会重新生成项目的README.md文件。
	2. helper n 会生成第n题的答题文件夹。
	3. helper -v 查看helper的版本`
)




var cfg config
var cfgFile = "leetcode.toml"
var lcFile = "leetcode.json"
var unavailableFile = "unavailable.json"

func init() {
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
	case "readme":
		rebuildReadme()
	default:
		buildProblemDir(os.Args[1])
	}
}
