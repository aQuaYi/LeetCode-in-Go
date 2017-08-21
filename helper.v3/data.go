package main

import (
	"fmt"
)

func updateData(categories []string) {
	for _, c := range categories {
		fmt.Println(getData(c))
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
