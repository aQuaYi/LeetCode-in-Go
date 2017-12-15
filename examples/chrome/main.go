package main

import (
	"os/exec"
)

func main() {

	// 执行单个shell命令时, 直接运行即可
	cmd := exec.Command("google-chrome", "www.baidu.com")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

}
