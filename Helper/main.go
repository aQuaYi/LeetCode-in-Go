package main

import (
	"log"
)

// 程序辅助设置
const (
	VERSION = "6.1.23"
)

func main() {
	log.Printf("Hi, %s. I'm %s\n", getConfig().Username, VERSION)

	cli := new(CLI)
	cli.Run()
}
