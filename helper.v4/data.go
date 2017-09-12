package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/aQuaYi/GoKit"
)

func update(categories []string) *leetcode {
	newLC := lastest(categories)
	oldLC, err := readLeetCodeRecord()
	if err != nil {
		log.Println("LeetCode 记录读取失败，无法与新记录对比:", err)
	} else {
		diff(newLC, oldLC)
	}

	saveLC(newLC)

	return newLC
}

func lastest(categories []string) *leetcode {
	lc := newLeetCode()
	for _, c := range categories {
		d := getData(c)
		lc.update(d)
	}

	lc.totalCategory()
	lc.getRanking()

	sort.Sort(lc.Problems)

	return lc
}

func readLeetCodeRecord() (*leetcode, error) {
	if !GoKit.Exist(lcFile) {
		msg := fmt.Sprintf("%s 不存在", lcFile)
		return nil, errors.New(msg)
	}

	raw := read(lcFile)
	lc := leetcode{}
	if err := json.Unmarshal(raw, &lc); err != nil {
		msg := fmt.Sprintf("获取 %s 失败：%s", lcFile, err)
		return nil, errors.New(msg)
	}

	return &lc, nil
}

func diff(new, old *leetcode) {
	// 对比 ranking
	nr, or := new.Ranking, old.Ranking
	if nr > 0 && or > 0 {
		v, delta := "进步了", or-nr
		if nr > or {
			v, delta = "后退了", nr-or
		}
		log.Printf("当前排名 %d，%s %d 名", nr, v, delta)
	} else {
		log.Printf("当前排名 %d", nr)
	}
	// 对比 已完成的问题
	lenNew := len(new.Problems)
	lenOld := len(old.Problems)
	isChanged := false

	i := 0
	for i < lenOld {
		n, o := new.Problems[i], old.Problems[i]

		if n.ID != o.ID {
			log.Println(n, o)
			log.Fatalln("LeetCode 的 Problems 数据出现错位。")
		}

		if n.IsAccepted == true && o.IsAccepted == false {
			log.Printf("新完成 %d %s", n.ID, n.Title)
			isChanged = true
		}

		i++
	}

	if !isChanged {
		log.Println("～ 没有新完成习题 ～")
	}

	for i < lenNew {
		log.Printf("新题 %d %s", new.Problems[i].ID, new.Problems[i].Title)
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

	log.Println("最新的 LeetCode 记录已经保存。")
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
