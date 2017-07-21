package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/aQuaYi/GoKit"
)

func (c *Category) run() Solveds {
	checkUser(c.User)

	checkDir(c.Name)

	c.compute()

	res := c.analysis()

	return res
}

// checkUser 检查是否下载了本人的数据
func checkUser(u string) {
	if u != USER {
		log.Fatalln("下载的不是本人的数据，请按照helper的说明文档，重新获取leetcode.cookie")
	}
}

// checkDir 检查是否存在分类目录
// 不存在的话，就新建一个
func checkDir(dir string) {
	if GoKit.Exist(dir) {
		return
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatalf("无法创建文件夹%s", dir)
	}

	log.Printf("已经创建文件夹%s", dir)
}

// 根据已有的内容，填上空缺字段
func (c *Category) compute() {
	for _, p := range c.Problems {
		// 统计各个难度的题目数量
		switch p.Difficulty.Level {
		case 1:
			c.Easy++
		case 2:
			c.Medium++
		case 3:
			c.Hard++
		default:
			log.Fatalln("出现了第4种难度")
		}

		p.PassRate = fmt.Sprintf("%d%%", 100*p.ACs/p.Submitted)
	}

	// 统计AC的总数
	c.AC = c.ACEasy + c.ACMedium + c.ACHard

	if c.Total != c.Easy+c.Medium+c.Hard {
		log.Fatalf("%s分类下的各难度题目数量之和，不等于Total", c.Name)
	}
}

// analysis 把已经解答的问题提取出来
func (c *Category) analysis() Solveds {
	res := Solveds{}

	for _, p := range c.Problems {
		pdir := p.checkDir(c.Name)
		// 检查处理已经AC的题目
		if p.Status == "ac" {
			res = append(res, makeSolved(*p, pdir))
		}
	}

	return res
}

// Category 保存种类信息
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
	Total    int        `json:"num_total"`
	Problems []*Problem `json:"stat_status_pairs"`
}

// Categories 收集各个种类的信息
type Categories []*Category

func (cs Categories) String() string {
	res := fmt.Sprintln("||Easy|Medium|Hard|Total|")
	res += fmt.Sprintln("|:--|:--:|:--:|:--:|:--:|")
	for _, c := range cs {
		res += fmt.Sprintln(c)
	}

	return res
}

func (c *Category) update(sub *Category) {
	c.ACEasy += sub.ACEasy
	c.ACMedium += sub.ACMedium
	c.ACHard += sub.ACHard
	c.AC += sub.AC

	c.Easy += sub.Easy
	c.Medium += sub.Medium
	c.Hard += sub.Hard
	c.Total += sub.Total
}

func (c *Category) String() string {
	res := fmt.Sprintf("|%s|", c.Name)
	res += fmt.Sprintf("%d/%d|", c.ACEasy, c.Easy)
	res += fmt.Sprintf("%d/%d|", c.ACMedium, c.Medium)
	res += fmt.Sprintf("%d/%d|", c.ACHard, c.Hard)
	res += fmt.Sprintf("%d/%d|", c.AC, c.Total)
	return res
}
