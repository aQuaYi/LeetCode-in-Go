package main

import (
	"fmt"
)

const (
	// USER 是你在leetcode.com的用户名
	// 登录后，在页面右上角可以看到
	USER = "aQuaYi"
)

func main() {
	categories := []string{
		"algorithms",
		"shell",
		"draft",
		"database",
		"system-design",
		"operating-system",
	}

	s := Solveds{}

	for i := range categories {
		j := getJSON(url(categories[i]))
		c := getCategory(j)
		s = append(s, c.run()...)
	}

	fmt.Println(s)
}
