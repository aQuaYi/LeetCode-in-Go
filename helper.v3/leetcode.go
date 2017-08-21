package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var cfg *config

func init() {
	if _, err := toml.DecodeFile(filename, cfg); err != nil {
		log.Fatalf(err.Error())
	}
}

type leetcode struct {
	Username string
	Ranking  int

	Categories []category
	Problems   []problem
}

// TODO: 去除 username 和 ranking 的全局变量
func newLeetCode() *leetcode {
	return &leetcode{
		Username: cfg.Login,
	}
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
		log.Fatalln("下载的非本人数据。程序可能没有signin。")
	}
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
