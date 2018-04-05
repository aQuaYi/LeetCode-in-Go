package main

import "log"

const (
	taskFile = "tasks.txt" // 任务输出文件夹名称
)

func makeTaskFile(pre string, first, last int) {
	log.Println("开始，生成新任务")
	res := ""
	count := 0
	ps := newLeetCode().Problems
	for i := first; i <= last && i < len(ps); i++ {
		if ps[i].ID == 0 {
			continue
		}

		res += ps[i].didaTask(pre) + "\n"
		count++
	}

	log.Printf("一共生成了 %d 条新任务\n", count)

	write(taskFile, res)
}

func collectReTask(ps problems, first, last int) string {
	res := ""
	count := 0
	for i := first; i <= last && i < len(ps); i++ {
		if ps[i].ID == 0 ||
			!ps[i].IsAccepted {
			continue
		}

		res += ps[i].didaTask("re") + "\n"
		count++
	}

	log.Printf("一共生成了 %d 条新任务\n", count)
	return res
}
