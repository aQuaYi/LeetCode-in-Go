package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func buildReadme() {
	log.Println("开始，重建 README 文档")

	lc := newLeetCode()
	makeReadmeFile(lc)

	// 放在 makefile 中实现这个功能了。
	// // 利用 chrome 打开github.com/aQuaYi/LeetCode-in-Go
	// link := "https://github.com/aQuaYi/LeetCode-in-Go#leetcode-%E7%9A%84-go-%E8%A7%A3%E7%AD%94"
	// cmd := exec.Command("google-chrome", link)
	// _, err := cmd.Output()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Println("正在打开项目主页")

	log.Println("完成，重建 README 文档")
}

func makeReadmeFile(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	var b bytes.Buffer

	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))

	err := tmpl.Execute(&b, lc)
	if err != nil {
		log.Fatal(err)
	}

	// 保存 README.md 文件
	write(file, string(b.Bytes()))
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
