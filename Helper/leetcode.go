package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
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
	if err != nil {
		log.Println("读取 LeetCode 的记录失败，正在重新生成 LeetCode 记录。失败原因：", err.Error())
		return getLeetCode()
	}
	return lc
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
	// 对比 ranking
	nr, or := new.Ranking, old.Ranking
	if nr > 0 && or > 0 {
		verb, delta := "进步", or-nr
		if nr > or {
			verb, delta = "后退", nr-or
		}
		log.Printf("当前排名 %d，%s了 %d 名", nr, verb, delta)
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
			log.Printf("～新完成～ %d.%s", n.ID, n.Title)
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
