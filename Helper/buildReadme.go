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

	makeMyFavoriteFile(lc)

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

func makeMyFavoriteFile(lc *leetcode) {
	file := "Favorite.md"
	os.Remove(file)

	var b bytes.Buffer

	tmpl := template.Must(template.New("favorite").Parse(readTMPL("favorite.markdown")))

	err := tmpl.Execute(&b, lc)
	if err != nil {
		log.Fatal(err)
	}

	// 保存 README.md 文件
	write(file, string(b.Bytes()))
}
