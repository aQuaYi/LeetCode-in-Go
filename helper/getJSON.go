package main

import (
	"fmt"
	"os/exec"
)

func getJSON(URL string) []byte {
	bytes, _ := exec.Command("curl", "-b", "leetcode.cookie", URL).Output()
	return bytes
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"

	return fmt.Sprintf(format, s)
}
