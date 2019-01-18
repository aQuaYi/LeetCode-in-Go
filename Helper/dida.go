package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	mail "gopkg.in/gomail.v2"
)

const (
	didaTaskFile = "dida.task.txt"
)

func dida(prefix string, p problem) {
	// 不再往滴答清单中添加任务
	// task := p.didaTask(prefix)
	// mailToDida(task)
}

func mailToDida(task string) {
	task += " ^LeetCode "
	task = delay(task)

	cfg := getConfig()
	if cfg.SMTP == "" || cfg.Port == 0 || cfg.EmailPassword == "" ||
		cfg.From == "" || cfg.To == "" {
		log.Printf("%v, 没有配置 Email，无法发送任务", cfg)
		saveLocal(task)
		return
	}

	m := mail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", cfg.To)
	m.SetHeader("Subject", task)
	m.SetBody("text/plain", fmt.Sprintf("添加日期 %s", time.Now()))
	d := mail.NewDialer(cfg.SMTP, cfg.Port, cfg.From, cfg.EmailPassword)

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
			log.Panicf("无法读取 %s：%s\n", didaTaskFile, err)
		}
		f, _ := os.Create(didaTaskFile)
		f.Close()
	}

	ts = append(ts, []byte(task+"\n")...)

	err = ioutil.WriteFile(didaTaskFile, ts, 0755)
	if err != nil {
		log.Panicf("无法写入 %s: %s\n", didaTaskFile, err)
	}

	log.Printf("新建任务已经写入 %s，请手动添加到滴答清单", didaTaskFile)
}

var m = map[string]time.Duration{
	"#do": 15,
	"#re": 90,
	"#fa": 30,
}

func delay(task string) string {
	key := task[:3]
	if day, ok := m[key]; ok {
		task += time.Now().Add(time.Hour * 24 * day).Format("2006-01-02")
		m[key] += 2
	}
	return task
}
