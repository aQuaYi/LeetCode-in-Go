package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/aQuaYi/GoKit"
)

const (
	cookieName = "leetcode.cookie"
)

func getCategory(name string) *Category {
	checkCookie(cookieName)

	URL := url(name)

	data := getJSON(URL)

	res := new(Category)
	if err := json.Unmarshal(data, res); err != nil {
		log.Fatal("无法把json转换成Category")
	}

	return res
}

func getJSON(URL string) []byte {
	log.Printf("开始下载%s的数据", URL)
	bytes, _ := exec.Command("curl", "-b", cookieName, URL).Output()
	return bytes
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"
	return fmt.Sprintf(format, s)
}

func checkCookie(name string) {
	if !GoKit.Exist(name) {
		log.Fatalf("%s 文件不存在，请按照helper中README.md中的说明，配置 %s", name, name)
	}
}
