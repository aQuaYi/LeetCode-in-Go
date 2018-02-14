package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func printTasks() {
	// 生成任务清单前，先重制 README 文件
	rebuildReadme()

	lc := lastestLeetCode()
	makeTasksFile(lc.Problems)
}

func makeTasksFile(problems problems) {
	content := ""
	for _, p := range problems {
		if !p.IsAccepted && p.IsAvailable {
			content = fmt.Sprintf("%d - #%d分 - %s - %s \n", p.ID, p.Difficulty, p.PassRate, p.Title) + content
		}
	}
	// 保存 taskFile 文件
	err := ioutil.WriteFile(tasksFile, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
