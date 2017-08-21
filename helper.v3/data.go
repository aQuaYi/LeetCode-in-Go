package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"

	"github.com/aQuaYi/GoKit"
)

func update(categories []string) {
	newLC := lastest(categories)
	oldLC := readFile()
	diff(newLC.Problems, oldLC.Problems)
	saveLC(newLC)
}

func lastest(categories []string) *leetcode {
	lc := newLeetCode()
	for _, c := range categories {
		d := getData(c)
		lc.update(d)
	}

	sort.Sort(lc.Problems)

	return lc
}

func readFile() *leetcode {
	lc := leetcode{}
	if !GoKit.Exist(lcFile) {
		log.Printf("%s 不存在", lcFile)
		return &lc
	}

	raw := read(lcFile)

	if err := json.Unmarshal(raw, &lc); err != nil {
		log.Fatalf("获取 %s 失败：%s", lcFile, err)
	}

	return &lc
}

func diff(new, old problems) {
	lenNew := len(new)
	lenOld := len(old)
	if lenNew == lenOld {
		log.Println("没有新完成习题")
	}
	i, j := 0, 0
	for i < lenNew && j < lenOld {
		if new[i].ID == old[i].ID {
			i++
			j++
			continue
		}

		log.Printf("新完成 %d %s", new[i].ID, new[i].Title)
		i++
	}
}

func saveLC(lc *leetcode) {
	raw, err := json.Marshal(lc)
	if err != nil {
		log.Fatal("无法把Leetcode数据转换成[]bytes: ", err)
	}
	if err = ioutil.WriteFile(lcFile, raw, 0666); err != nil {
		log.Fatal("无法把Marshal后的lc保存到文件: ", err)
	}
}

// data 保存API信息
type data struct {
	Name     string          `json:"category_slug"`
	User     string          `json:"user_name"`
	ACEasy   int             `json:"ac_easy"`
	ACMedium int             `json:"ac_medium"`
	ACHard   int             `json:"ac_hard"`
	AC       int             `json:"num_solved"`
	Problems []problemStatus `json:"stat_status_pairs"`
}

type problemStatus struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	IsFavor    bool `json:"is_favor"`
	IsPaid     bool `json:"paid_only"`
	Difficulty `json:"difficulty"`
}

// State 保存单个问题的解答状态
type State struct {
	Title     string `json:"question__title"`
	TitleSlug string `json:"question__title_slug"`
	IsNew     bool   `json:"is_new_question"`
	ID        int    `json:"question_id"`
	ACs       int    `json:"total_acs"`
	Submitted int    `json:"total_submitted"`
}

// Difficulty 问题的难度
type Difficulty struct {
	Level int `json:"level"`
}
