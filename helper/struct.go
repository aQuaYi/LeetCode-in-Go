package main

import (
	"log"

	"os"

	"github.com/aQuaYi/GoKit"
)

type Category struct {
	Name     string `json:"category_slug"`
	User     string `json:"user_name"`
	ACEasy   int    `json:"ac_easy"`
	ACMedium int    `json:"ac_medium"`
	ACHard   int    `json:"ac_hard"`
	AC       int    `json:"num_solved"`
	Easy     int
	Medium   int
	Hard     int
	Total    int       `json:"num_total"`
	Problems []Problem `json:"stat_status_pairs"`
}

type Problem struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	PaidOnly   bool `json:"paid_only"`
	Difficulty `json:"difficulty`
}

type State struct {
	Title     string `json:"question__title"`
	TitleSlug string `json:"question__title_slug"`
	IsNew     bool   `json:"is_new_question"`
	ID        int    `json:"question_id"`
	ACs       int    `json:"total_acs"`
	Submitted int    `json:"total_submitted"`
	PassRate  float64
}

type Difficulty struct {
	Level int `json:"level"`
}

func (d *Category) run() {
	checkUser(d.User)
	checkDir(d.Name)

}

func checkUser(u string) {
	if u != USER {
		log.Fatalln("下载的不是本人的数据，请按照helper的说明文档，重新获取leetcode.cookie")
	}
}

func checkDir(dir string) {
	if GoKit.Exist(dir) {
		return
	}

	err := os.Mkdir(dir, 755)
	if err != nil {
		log.Fatalf("无法创建文件夹%s", dir)
	}

	log.Printf("已经创建文件夹%s", dir)
}
