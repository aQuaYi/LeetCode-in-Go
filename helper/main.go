package main

import (
	"fmt"
)
const (
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

	j := getJSON(url(categories[1]))

	fmt.Println(string(j))
}
