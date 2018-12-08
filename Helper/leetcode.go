package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	leetCodeJSON = "leetcode.json"
)

type leetcode struct {
	Username string    // 用户名
	Ranking  int       // 网站排名
	Updated  time.Time // 数据更新时间

	Record   record   // 已解答题目与全部题目的数量，按照难度统计
	Problems problems // 所有问题的集合
}

func newLeetCode() *leetcode {
	log.Println("开始，获取 LeetCode 数据")

	lc, err := readLeetCode()
	if err != nil {
		log.Println("读取 LeetCode 的记录失败，正在重新生成 LeetCode 记录。失败原因：", err.Error())
		lc = getLeetCode()
	}

	lc.refresh()

	lc.save()

	log.Println("获取 LeetCode 的最新数据")
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

	if err := os.Remove(leetCodeJSON); err != nil {
		log.Panicf("删除 %s 失败，原因是：%s", leetCodeJSON, err)
	}

	raw, err := json.MarshalIndent(lc, "", "\t")
	if err != nil {
		log.Fatal("无法把Leetcode数据转换成[]bytes: ", err)
	}
	if err = ioutil.WriteFile(leetCodeJSON, raw, 0666); err != nil {
		log.Fatal("无法把 Marshal 后的 lc 保存到文件: ", err)
	}
	log.Println("最新的 LeetCode 记录已经保存。")
	return
}

func (lc *leetcode) refresh() {
	if time.Since(lc.Updated) < time.Minute {
		log.Printf("LeetCode 数据在 %s 前刚刚更新过，跳过此次刷新\n", time.Since(lc.Updated))
		return
	}

	log.Println("开始，刷新 LeetCode 数据")
	newLC := getLeetCode()

	logDiff(lc, newLC)

	*lc = *newLC
}

func logDiff(old, new *leetcode) {
	// 对比 ranking
	str := fmt.Sprintf("当前排名 %d", new.Ranking)
	verb, delta := "进步", old.Ranking-new.Ranking
	if new.Ranking > old.Ranking {
		verb, delta = "后退", new.Ranking-old.Ranking
	}
	str += fmt.Sprintf("，%s了 %d 名", verb, delta)
	log.Println(str)

	lenOld, lenNew := len(old.Problems), len(new.Problems)
	hasNewFinished := false

	i := 0

	// 检查新旧都有的问题
	for i < lenOld && i < lenNew {
		o, n := old.Problems[i], new.Problems[i]
		// 检查是 n 是否是新 完成
		if o.IsAccepted == false && n.IsAccepted == true {
			log.Printf("～新完成～ %d - %s", n.ID, n.Title)
			dida("re", n)
			hasNewFinished = true
		}
		// 检查是 n 是否是新 收藏
		if o.IsFavor == false && n.IsFavor == true {
			log.Printf("～新收藏～ %d - %s", n.ID, n.Title)
			dida("fa", n)
		} else if o.IsFavor == true && n.IsFavor == false {
			log.Printf("～取消收藏～ %d - %s", o.ID, o.Title)
			time.Sleep(time.Second)
		}

		// 有时候，会在中间添加新题
		if o.Title == "" && n.Title != "" {
			log.Printf("新题: %d - %s", new.Problems[i].ID, new.Problems[i].Title)
			dida("do", n)
		}

		i++
	}

	log.Printf("已经检查完了 %d 题\n", i)

	if !hasNewFinished {
		log.Println("～ 没有新完成习题 ～")
	}

	// 检查新添加的习题
	for i < lenNew {
		if new.Problems[i].isAvailable() {
			log.Printf("新题: %d - %s", new.Problems[i].ID, new.Problems[i].Title)
			dida("do", new.Problems[i])
		}
		i++
	}
}

func (lc *leetcode) ProgressTable() string {
	return lc.Record.progressTable()
}

func (lc *leetcode) AvailableTable() string {
	return lc.Problems.available().table()
}

func (lc *leetcode) FavoriteTable() string {
	return lc.Problems.favorite().table()
}

func (lc *leetcode) FavoriteCount() int {
	return len(lc.Problems.favorite())
}

func (lc *leetcode) UnavailableList() string {
	res := lc.Problems.unavailable().list()
	// 为了 README.md 文档的美观，需要删除最后一个换行符号
	return res[:len(res)-1]
}
