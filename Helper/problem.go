package main

import (
	"fmt"
	"strings"
)

type problem struct {
	ID                                 int
	Title                              string
	TitleSlug                          string
	PassRate                           string
	Difficulty                         string
	IsAccepted, IsPaid, IsFavor, IsNew bool
	HasNoGoOption                      bool // 不能够使用 Go 语言解答
}

func newProblem(ps problemStatus) problem {
	level := []string{"", "Easy", "Medium", "Hard"}

	p := problem{
		ID:        ps.State.ID,
		Title:     ps.State.Title,
		TitleSlug: ps.State.TitleSlug,
		// p.Submitted + 1 是因为刚刚添加的新题的 submitted 为 0
		PassRate:   fmt.Sprintf("%d%%", ps.ACs*100/(ps.Submitted+1)),
		Difficulty: level[ps.Difficulty.Level],
		IsAccepted: ps.Status == "ac",
		IsPaid:     ps.IsPaid,
		IsFavor:    ps.IsFavor,
		IsNew:      ps.State.IsNew,
	}

	return p
}

func (p problem) isAvailable() bool {
	if p.ID == 0 || p.IsPaid || p.HasNoGoOption {
		return false
	}
	return true
}

func (p problem) Dir() string {
	path := "Algorithms"
	return fmt.Sprintf("./%s/%04d.%s", path, p.ID, p.TitleSlug)
}

func (p problem) link() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
}

func (p problem) tableLine() string {
	// 题号
	res := fmt.Sprintf("|[%04d](%s)|", p.ID, p.link())

	// 标题
	t := ""
	if p.IsAccepted {
		t = fmt.Sprintf(`[%s](%s)`, strings.TrimSpace(p.Title), p.Dir())
	} else {
		t = fmt.Sprintf(` * %s`, p.Title)
	}
	if p.IsNew {
		t += " :new: "
	}
	res += t + "|"

	// 通过率
	res += fmt.Sprintf("%s|", p.PassRate)

	// 难度
	res += fmt.Sprintf("%s|", p.Difficulty)

	// 收藏
	f := ""
	if p.IsFavor {
		f = "[❤](https://leetcode.com/list/oussv5j)"
	}
	res += fmt.Sprintf("%s|\n", f)

	return res
}

func (p problem) listLine() string {
	return fmt.Sprintf("- [%d.%s](%s)\n", p.ID, p.Title, p.link())
}

func (p problem) didaTask(prefix string) string {
	return fmt.Sprintf("#%s - %04d - #%s - %s - %s - %s", prefix, p.ID, p.Difficulty, p.PassRate, p.Title, p.link())
}
