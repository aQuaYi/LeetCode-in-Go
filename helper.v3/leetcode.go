package main

import (
	"fmt"
	"log"
	"sort"
)

type leetcode struct {
	Username string
	Ranking  int

	Categories []category
	Problems   problems
}

// TODO: 去除 username 和 ranking 的全局变量
func newLeetCode() *leetcode {
	return &leetcode{
		Username: cfg.Login,
	}
}
func (l *leetcode) totalCategory() {
	t := category{
		Name: "Total",
	}

	for _, c := range l.Categories {
		t.Easy.Solved += c.Easy.Solved
		t.Easy.Total += c.Easy.Total
		t.Medium.Solved += c.Medium.Solved
		t.Medium.Total += c.Medium.Total
		t.Hard.Solved += c.Hard.Solved
		t.Hard.Total += c.Hard.Total
		t.Total.Solved += c.Total.Solved
		t.Total.Total += c.Total.Total
	}

	l.Categories = append(l.Categories, t)
}
func (l *leetcode) update(d *data) {
	l.check(d)
	ps, e, m, h := countData(d)
	l.Problems = append(l.Problems, ps...)
	c := newCategory(d, e, m, h)
	l.Categories = append(l.Categories, c)
}

func (l *leetcode) check(d *data) {
	if d.User != l.Username {
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
			ID:         p.ID,
			Title:      p.Title,
			TitleSlug:  p.TitleSlug,
			PassRate:   fmt.Sprintf("%d%%", p.ACs*100/p.Submitted),
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

func newCategory(d *data, e, m, h int) category {
	c := category{
		Name: d.Name,
	}

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

func (ps problems) String() string {
	sort.Sort(ps)
	res := "|题号|题目|难度|总体通过率|收藏|\n"
	res += "|:-:|:-|:-: | :-: | :-: |\n"
	for _, p := range ps {
		res += fmt.Sprintln(p)
	}
	return res
}

type problem struct {
	ID                         int
	Title, TitleSlug           string
	PassRate                   string
	Difficulty                 int
	IsAccepted, IsFavor, IsNew bool
}

func (p problem) page() string {
	format := "https://leetcode.com/problems/%s"
	return fmt.Sprintf(format, p.TitleSlug)
}

// TODO:  获取ranking
// // 获取 ranking
// temp := getRanking(username)
// var err error
// ranking, err = strconv.Atoi(temp)
// if err != nil {
// 	log.Fatalf("获取的 %s 无法转换成 ranking", temp)
// }

// log.Printf("%s, your ranking is %d\n", username, ranking)
