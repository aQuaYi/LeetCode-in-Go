package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func getLeetCode() *leetcode {

	return nil
}

func getCategoryData(name string) *data {
	URL := url(name)

	raw := getRaw(URL)

	res := new(data)
	if err := json.Unmarshal(raw, res); err != nil {
		log.Fatal("无法把json转换成Category: " + err.Error())
	}

	return res
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"
	return fmt.Sprintf(format, s)
}
