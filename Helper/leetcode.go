package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"time"
)

const (
	leetCodeJSON = "leetcode.json"
)

type leetcode struct {
	Username string

	Record   record   // 已解答题目与全部题目的数量，按照难度统计
	Problems problems // 所有问题的集合

	Ranking int

	updated time.Time
}

func newLeetCode() *leetcode {
	lc, err := readLeetCode()
	if err == nil {
		return lc
	}

	log.Println("读取 LeetCode 的记录失败，正在重新生成 LeetCode 记录。失败原因：", err.Error())
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

func (lc *leetcode) update() {
	if time.Since(lc.updated) < 5*time.Minute {
		log.Printf("数据刚刚在 %s 更新过，跳过此次更新\n", lc.updated)
		return
	}

	newLC := getLeetCode()
	logDiff(lc, newLC)
	lc = newLC
}

func logDiff(old, new *leetcode) {
	log.Fatal("logDiff is vaild")
	return
}
