package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configTOML = "config.toml"
)

type config struct {
	Username string
	Password string
	Cookie string

	// 以下是电子邮件设置
	SMTP          string
	Port          int
	From          string
	To            string
	EmailPassword string
}

func (c config) String() string {
	format := "Username: %s, Password: %s, SMTP: %s, Port: %d, From: %s, To: %s, EmailPassword: %s "
	return fmt.Sprintf(format,
		c.Username,
		c.Password,
		c.SMTP,
		c.Port,
		c.From,
		c.To,
		c.EmailPassword)
}

func getConfig() *config {
	cfg := new(config)

	if _, err := toml.DecodeFile(configTOML, &cfg); err != nil {
		log.Panicf(err.Error())
	}

	// log.Printf("get config: %s", cfg)

	return cfg
}
