package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func printTasks() {
	lc := lastestLeetCode()
	makeTasksFile(lc.Problems)
}

func makeTasksFile(problems problems) {
	content := ""
	for _, p := range problems {
		content += fmt.Sprintf("%d - %s%% - %d - %s\n", p.Difficulty, p.PassRate, p.ID, p.Title)
	}
	// 保存 taskFile 文件
	err := ioutil.WriteFile(tasksFile, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
