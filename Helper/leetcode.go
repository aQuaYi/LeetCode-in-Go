package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

const (
	leetCodeJSON = "leetcode.json"
)

type leetcode struct {
	Username string

	// Problems problems // 所有问题的集合
	// Recode   recode   // 已解答题目与全部题目的数量，按照难度统计

	Ranking int
}

func newLeetCode() *leetcode {
	lc, err := readLeetCode()
	if err == nil {
		return lc
	}

	log.Println("读取 LeetCode 的记录失败，新生成一个记录。失败原因：", err.Error())
	return getLeetCode()
}

func readLeetCode() (*leetcode, error) {
	data, err := ioutil.ReadFile(leetCodeJSON)
	if err != nil {
		return nil, errors.New("读取文件失败：" + err.Error())
	}

	lc := new(leetcode)
	if err := json.Unmarshal(data, lc); err != nil {
		return nil, errors.New("转换成 leetcode 时，失败：" + err.Error())
	}

	return lc, nil
}

func (lc *leetcode) save() {
	raw, err := json.Marshal(lc)
	if err != nil {
		log.Fatal("无法把Leetcode数据转换成[]bytes: ", err)
	}
	if err = ioutil.WriteFile(leetCodeJSON, raw, 0666); err != nil {
		log.Fatal("无法把Marshal后的lc保存到文件: ", err)
	}
	log.Println("最新的 LeetCode 记录已经保存。")
	return
}
