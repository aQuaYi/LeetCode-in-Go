package main

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

func didaTask(task string) {
	cfg := getConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To)
	m.SetHeader("Subject", task)

	d := gomail.NewDialer(cfg.SMTP, cfg.Port, cfg.From, cfg.EmailPasswd)

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("无法发送任务到 滴答清单：", err)
	}
}
