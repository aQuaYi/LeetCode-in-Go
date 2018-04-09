package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	gomail "gopkg.in/gomail.v2"
)

const (
	didaTaskFile = "dida.task"
)

func dida(prefix string, p problem) {
	task := p.didaTask(prefix)
	mailToDida(task)
}

func mailToDida(task string) {
	cfg := getConfig()

	if cfg.SMTP == "" || cfg.Port == 0 || cfg.EmailPasswd == "" ||
		cfg.From == "" || cfg.To == "" {
		log.Println("没有配置 Email，无法发送任务")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To)
	m.SetHeader("Subject", task+" ^LeetCode")
	m.SetBody("text/plain", fmt.Sprintf("添加日期 %s", time.Now()))

	d := gomail.NewDialer(cfg.SMTP, cfg.Port, cfg.From, cfg.EmailPasswd)

	if err := d.DialAndSend(m); err != nil {
		log.Println("无法发送任务到 滴答清单：", err)
		saveLocal(task)
		return
	}

	log.Printf("已经在滴答清单中添加任务： %s", task)
}

func saveLocal(task string) {
	ts, err := ioutil.ReadFile(didaTaskFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("无法读取 %s：%s\n", didaTaskFile, err)
		}
		os.Create(didaTaskFile)
	}

	ts = append(ts, []byte(task+"\n")...)

	err = ioutil.WriteFile(didaTaskFile, ts, 0755)
	if err != nil {
		log.Fatalf("无法写入 %s: %s\n", didaTaskFile, err)
	}

	log.Printf("新建任务已经写入 %s，请手动添加到滴答清单", didaTaskFile)
}
