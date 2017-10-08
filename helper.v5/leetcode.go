package main

import (
	"fmt"
	"log"
	"sort"
)

type leetcode struct {
	Username string

	Problems   problems
	Algorithms category

	Progress int

	Ranking int
}

func newLeetCode() *leetcode {
	lc := &leetcode{}

	lc.Username = cfg.Login

	lc.Problems, lc.Algorithms = newAlgorithms()

	lc.Ranking = getRanking(lc.Username)

	return lc
}

func newAlgorithms() (problems, category) {
	d := getData("Algorithms")
	check(d)
	ps, e, m, h := countData(d)
	c := newCategory(d, e, m, h)
	return ps, c
}

func check(d *data) {
	if d.User != cfg.Login {
		log.Fatalln("下载了非本人的数据。")
	}
	log.Printf("%s 通过检查", d.Name)
}

func countData(d *data) (ps []problem, e, m, h int) {
	for _, p := range d.Problems {

		if p.IsPaid {
			continue
		}

		temp := problem{
			ID:        p.ID,
			Dir:       fmt.Sprintf("./%s/%04d.%s", d.Name, p.ID, p.TitleSlug),
			Title:     p.Title,
			TitleSlug: p.TitleSlug,
			// p.Submitted + 1 是因为刚刚添加的新题的 submitted 为 0
			PassRate:   fmt.Sprintf("%d%%", p.ACs*100/(p.Submitted+1)),
			Difficulty: p.Difficulty.Level,
			IsAccepted: p.Status == "ac",
			IsFavor:    p.IsFavor,
			IsNew:      p.IsNew,
		}
		ps = append(ps, temp)
		switch temp.Difficulty {
		case 1:
			e++
		case 2:
			m++
		case 3:
			h++
		default:
			log.Fatalln("题目出现了第4种难度", p.ID, p.Title)
		}
	}
	return
}

type category struct {
	Name                      string
	Easy, Medium, Hard, Total count
}

func (c category) progressTable() string {
	res := fmt.Sprintln("|Easy|Medium|Hard|Total|")
	res += fmt.Sprintln("|:---:|:---:|:---:|:---:|")
	res += fmt.Sprintf("%d / %d|", c.Easy.Solved, c.Easy.Total)
	res += fmt.Sprintf("%d / %d|", c.Medium.Solved, c.Medium.Total)
	res += fmt.Sprintf("%d / %d|", c.Hard.Solved, c.Hard.Total)
	res += fmt.Sprintf("%d / %d|", c.Total.Solved, c.Total.Total)
	return res
}

func newCategory(d *data, e, m, h int) category {
	c := category{}

	c.Easy.Solved = d.ACEasy
	c.Easy.Total = e

	c.Medium.Solved = d.ACMedium
	c.Medium.Total = m

	c.Hard.Solved = d.ACHard
	c.Hard.Total = h

	c.Total.Solved = d.ACEasy + d.ACMedium + d.ACHard
	c.Total.Total = e + m + h

	return c
}

type count struct {
	Solved, Total int
}

type problems []problem

func (ps problems) Len() int {
	return len(ps)
}

func (ps problems) Less(i, j int) bool {
	return ps[i].ID < ps[j].ID
}

func (ps problems) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps problems) accepted() problems {
	res := make([]problem, 0, len(ps))
	for _, p := range ps {
		if p.IsAccepted {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) table() string {
	sort.Sort(ps)
	res := "|题号|题目|难度|总体通过率|收藏|\n"
	res += "|:-:|:-|:-: | :-: | :-: |\n"
	for _, p := range ps {
		res += p.tableLine()
	}
	return res
}

func (ps problems) list() string {
	sort.Sort(ps)
	res := ""
	for _, p := range ps {
		res += p.listLine()
	}
	return res
}

type problem struct {
	ID                         int
	Dir                        string
	Title, TitleSlug           string
	PassRate                   string
	Difficulty                 int
	IsAccepted, IsFavor, IsNew bool
}

func (p problem) link() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
}

func (p problem) tableLine() string {
	res := fmt.Sprintf("|%d|", p.ID)
	res += fmt.Sprintf(`[%s](%s)|`, p.Title, p.Dir)
	res += fmt.Sprintf("%s|", degrees[p.Difficulty])
	res += fmt.Sprintf("%s|", p.PassRate)
	f := " "
	if p.IsFavor {
		f = "❤"
	}
	res += fmt.Sprintf("%s|", f)
	return res
}

var degrees = map[int]string{
	1: `☆`,
	2: `☆ ☆`,
	3: `☆ ☆ ☆`,
}

func (p problem) listLine() string {
	return fmt.Sprintf("[%d. %s](%s)", p.ID, p.Title, p.link())
}
