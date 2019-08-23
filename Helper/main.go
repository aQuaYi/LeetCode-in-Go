package main

import (
	"log"
)

// 程序辅助设置
const (
	VERSION = "7.0.8"
)

func main() {
	log.Printf("Hi, %s. I'm %s\n", getConfig().Username, VERSION)

	cli := new(CLI)
	cli.Run()
}
