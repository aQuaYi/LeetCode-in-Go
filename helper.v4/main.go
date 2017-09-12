package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

// 程序辅助设置
const (
	VERSION = "4.0.0"
	USAGE   = `使用方法：
	1. 运行 helper 会重新生成项目的README.md。
	2. 运行 helper 123 会生成第123题的答题文件夹。`
)

var problemNum int
var cfg config
var cfgFile = "leetcode.toml"
var lcFile = "leetcode.json"

func init() {
	// 启动时，导入配置
	if _, err := toml.DecodeFile(cfgFile, &cfg); err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Hi, %s. \n", cfg.Login)

	// 对 req 赋值
	req = request.NewRequest(new(http.Client))
}

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

	// 下载题目资料不需要登录 leetcode
	if problemNum > 0 {
		lc, err := readLeetCodeRecord()
		if err != nil {
			log.Fatalln("读取 LeetCode 记录失败: ", err)
		}
		makeProblemDir(lc.Problems, problemNum)
		return
	}

	signin()

	categories := []string{
		"Algorithms",
		// "database",
		"Draft",
		// "operating-system",
		// "shell",
		// "system-design",
	}

	lc := update(categories)

	makeREADME(lc)

}
