package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseFunction(fc string) (fcName, para, ansType, nfc string) {
	log.Println("正在分解函数")

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fcName = "myFunc"
			para = "p int"
			ansType = "int"
			nfc = "func myFunc(p int) int {\n\nreturn nil\n}"
		}
	}()

	funcIndex := strings.Index(fc, "func ")
	a := funcIndex + strings.Index(fc[funcIndex:], " ")
	b := funcIndex + strings.Index(fc[funcIndex:], "(")
	c := funcIndex + strings.Index(fc[funcIndex:], ")")
	d := funcIndex + strings.Index(fc[funcIndex:], "{")

	fcName = fc[a+1 : b]
	para = fc[b+1 : c]
	ansType = strings.TrimSpace(fc[c+1 : d])
	nfc = fmt.Sprintf("func %s(%s) %s {\n\nreturn nil\n}", fcName, para, ansType)

	return
}

func getFunction(url string) string {
	inputReader := bufio.NewReader(os.Stdin)
	var err error
	fc := ""
	for !strings.HasPrefix(fc, "func ") {
		fmt.Print("请输入 Go 函数:")
		fc, err = inputReader.ReadString('\n')
		if err != nil {
			log.Panicf("读取 Go 函数失败：%s", err)
		}
	}

	return fc

}
