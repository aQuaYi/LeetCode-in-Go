package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configTOML = "config.toml"
)

type config struct {
	Username string
	Password string

	// 以下是电子邮件设置
	SMTP   string
	Port   int
	From   string
	To     string
	Passwd string
}

func getConfig() *config {
	cfg := new(config)

	if _, err := toml.DecodeFile(configTOML, &cfg); err != nil {
		log.Fatalf(err.Error())
	}

	return cfg
}
