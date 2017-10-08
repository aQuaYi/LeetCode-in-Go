package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aQuaYi/GoKit"
)

func lastestLeetCode() *leetcode {
	newLC := newLeetCode()

	showChange(newLC)

	saveLC(newLC)

	return newLC
}

func showChange(new *leetcode) {
	// 获取 oldLC
	old, err := readLeetCodeRecord()
	if err != nil {
		log.Println("LeetCode 记录读取失败，无法与新记录对比:", err)
		return
	}

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
			os.Remove("leetcode.json")
			log.Fatalln("LeetCode 的 Problems 数据出现错位。已经删除 leetcode.json。 请重试")
		}

		if n.IsAccepted == true && o.IsAccepted == false {
			log.Printf("新完成 %d.%s", n.ID, n.Title)
			isChanged = true
		}

		i++
	}

	if !isChanged {
		log.Println("～ 没有新完成习题 ～")
	}

	for i < lenNew {
		log.Printf("新题: %d.%s", new.Problems[i].ID, new.Problems[i].Title)
		i++
	}
}

func readLeetCodeRecord() (*leetcode, error) {
	if !GoKit.Exist(leetCodeFile) {
		msg := fmt.Sprintf("%s 不存在", leetCodeFile)
		return nil, errors.New(msg)
	}

	raw := read(leetCodeFile)
	lc := leetcode{}
	if err := json.Unmarshal(raw, &lc); err != nil {
		msg := fmt.Sprintf("获取 %s 失败：%s", leetCodeFile, err)
		return nil, errors.New(msg)
	}

	return &lc, nil
}

func saveLC(lc *leetcode) {
	raw, err := json.Marshal(lc)
	if err != nil {
		log.Fatal("无法把Leetcode数据转换成[]bytes: ", err)
	}
	if err = ioutil.WriteFile(leetCodeFile, raw, 0666); err != nil {
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
