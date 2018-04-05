package main

import (
	"fmt"
	"log"
	"time"

	gomail "gopkg.in/gomail.v2"
)

func dida(prefix string, p problem) {
	task := p.didaTask(prefix) + " ^LeetCode"
	mailToDida(task)
}

func mailToDida(task string) {
	cfg := getConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To)
	m.SetHeader("Subject", task)
	m.SetBody("text/html", fmt.Sprintf("添加日期 %s", time.Now()))

	d := gomail.NewDialer(cfg.SMTP, cfg.Port, cfg.From, cfg.EmailPasswd)

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("无法发送任务到 滴答清单：", err)
	}
}
