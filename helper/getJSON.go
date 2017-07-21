package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func getCategory(data []byte) *Category {
	res := new(Category)
	if err := json.Unmarshal(data, res); err != nil {
		log.Fatal("无法把json转换成Category")
	}
	return res
}
func getJSON(URL string) []byte {
	log.Printf("开始下载%s的数据", URL)
	bytes, _ := exec.Command("curl", "-b", "leetcode.cookie", URL).Output()
	return bytes
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"
	return fmt.Sprintf(format, s)
}
