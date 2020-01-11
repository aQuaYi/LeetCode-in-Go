package main

import (
	"log"
	"net/http"

	"github.com/mozillazg/request"
)

const (
	loginPageURL = "https://leetcode.com/accounts/login/"
)

var req *request.Request

func newReq() *request.Request {
	if req == nil {
		req = signin()
	}
	return req
}

// 登录 leetcode
// 返回的 req 带有 cookie
func signin() *request.Request {
	// log.Println("正在登录中...")
	cfg := getConfig()

	// 对 req 赋值
	req := request.NewRequest(new(http.Client))

	// 配置request
	req.Headers = map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": "",
		"cookie":          cfg.Cookie,
		"Referer":         "https://leetcode.com/accounts/login/",
		"origin":          "https://leetcode.com",
	}

	// login
	// csrfToken := getCSRFToken(req)

	// log.Printf("csrfToken: %s", csrfToken)

	// req.Data = map[string]string{
	// 	"csrfmiddlewaretoken": csrfToken,
	// 	"login":               cfg.Username,
	// 	"password":            cfg.Password,
	// 	"next":                "problems",
	// }
	// if err := login(req); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("成功登录")

	return req
}

func getCSRFToken(req *request.Request) string {
	resp, err := req.Get(loginPageURL)
	if err != nil {
		log.Panicf("无法 Get 到 %s: %s", loginPageURL, err)
	}

	cookies := resp.Cookies()

	for _, ck := range cookies {
		if ck.Name == "csrftoken" {
			return ck.Value
		}
	}

	panic("无法在 Cookies 中找到 csrftoken")
}

func login(req *request.Request) error {
	resp, err := req.Post(loginPageURL)
	defer resp.Body.Close() // **Don't forget close the response body**
	return err
}
