package main

import (
	"log"
)

const (
	taskFile = "tasks.txt" // 任务输出文件夹名称
)

func makeTaskFile(prefix string, first, last int) {
	log.Println("开始，生成新任务")

	res := ""
	count := 0
	ps := newLeetCode().Problems
	var isWanted func(int) bool
	collect := func() {
		for i := first; i <= last && i < len(ps); i++ {
			if !isWanted(i) {
				continue
			}
			res += ps[i].didaTask(prefix) + "\n"
			count++
		}
	}

	// 根据 prefix 的不同，设置不同的 isWanted
	switch prefix {
	case "do": // to do
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && !ps[i].IsAccepted && !ps[i].IsPaid && !ps[i].HasNoGoOption
		}
	case "re": // review
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsAccepted && !ps[i].IsPaid
		}
	case "mi": // miss
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsAccepted && ps[i].HasNoGoOption
		}
	case "fa": // favor
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsFavor
		}
	default:
		// 和 do 的一样
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && !ps[i].IsAccepted && !ps[i].IsPaid && !ps[i].HasNoGoOption
		}
	}

	collect()
	log.Printf("完成，一共生成了 %d 条新任务\n", count)
	write(taskFile, res)
}
