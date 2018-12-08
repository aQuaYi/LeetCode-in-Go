package main

import (
	"flag"
	"fmt"
	"log"

	"os"
)

const (
	// USAGE 说明了程序的使用方式
	USAGE = `使用方法：
	helper readme
		重新生成项目的 README.md 文件
	helper prepare -number N
		生成第 N 题的答题文件夹
	helper task -prefix PRE -first FIRST -last LAST
		对于所有 [FIRST, LAST] 之内的题目，生成类似
		PRE - 770 - #hard - 43% - Basic Calculator IV
		的条目，并保存到文件 tasks.txt 中
`
)

// CLI 负责处理命令行参数
type CLI struct{}

func (c *CLI) printUsage() {
	fmt.Println(USAGE)
}

func (c *CLI) checkArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (c *CLI) Run() {
	c.checkArgs()

	readmeCmd := flag.NewFlagSet("readme", flag.ExitOnError)

	prepareCmd := flag.NewFlagSet("prepare", flag.ExitOnError)
	problemNumber := prepareCmd.Int("number", 0, "请输入你想要准备的题目的题号")

	taskCmd := flag.NewFlagSet("task", flag.ExitOnError)
	prefix := taskCmd.String("prefix", "", "答题任务的前缀")
	first := taskCmd.Int("first", 0, "答题任务的最小题号")
	last := taskCmd.Int("last", 0, "答题任务的最大题号")

	switch os.Args[1] {
	case "readme":
		err := readmeCmd.Parse(os.Args[2:])
		check(err)
		buildReadme()
	case "prepare":
		err := prepareCmd.Parse(os.Args[2:])
		check(err)
		buildProblemDir(*problemNumber)
	case "task":
		err := taskCmd.Parse(os.Args[2:])
		check(err)
		makeTaskFile(*prefix, *first, *last)
	default:
		c.printUsage()
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
