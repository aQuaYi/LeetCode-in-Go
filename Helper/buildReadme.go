package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func buildReadme() {
	log.Println("开始，重建 README 文档")

	lc := newLeetCode()
	makeReadmeFile(lc)

	log.Println("完成，重建 README 文档")
}

func makeReadmeFile(lc *leetcode) {
	file := "README.md"
	os.Remove(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))

	tmpl.Execute(writer, lc)

	// 保存 README.md 文件
	write(file, string(b.Bytes()))
}

func getHead(lc *leetcode) string {
	headFormat := string(read("README_HEAD.md"))
	return fmt.Sprintf(headFormat, lc.Username, lc.Ranking, lc.Username)
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
