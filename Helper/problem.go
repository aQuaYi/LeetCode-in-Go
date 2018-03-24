package main

import (
	"fmt"
	"strings"
)

type problem struct {
	ID                                      int
	Dir                                     string
	Title, TitleSlug                        string
	PassRate                                string
	Difficulty                              int
	IsAccepted, IsFavor, IsNew, IsAvailable bool
}

func (p problem) link() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
}

func (p problem) tableLine() string {
	res := fmt.Sprintf("|%d|", p.ID)
	if p.IsAccepted {
		res += fmt.Sprintf(`[%s](%s)|`, strings.TrimSpace(p.Title), p.Dir)
	} else {
		res += fmt.Sprintf(` * %s|`, p.Title)
	}
	res += fmt.Sprintf("%s|", p.PassRate)
	res += fmt.Sprintf("%s|", degrees[p.Difficulty])
	f := " "
	if p.IsFavor {
		f = "[❤](https://leetcode.com/list/oussv5j)"
	}
	res += fmt.Sprintf("%s|\n", f)
	return res
}

var degrees = map[int]string{
	1: "Easy",
	2: "Medium",
	3: "Hard",
}

func (p problem) listLine() string {
	return fmt.Sprintf("- [%d.%s](%s)\n", p.ID, p.Title, p.link())
}
