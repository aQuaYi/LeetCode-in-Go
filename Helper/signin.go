package main

import (
	"log"
	"net/http"
	"regexp"

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
	log.Println("正在登录中...")
	cfg := getConfig()

	// 对 req 赋值
	req := request.NewRequest(new(http.Client))

	// 配置request
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         "https://leetcode.com/",
	}

	// login
	csrfToken, err := getCSRFToken(req)
	if err != nil {
		log.Fatal(err)
	}
	req.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"login":               cfg.Username,
		"password":            cfg.Password,
	}
	if err = login(req); err != nil {
		log.Fatal(err)
	}

	log.Println("成功登录")

	return req
}

func getCSRFToken(req *request.Request) (string, error) {
	resp, err := req.Get(loginPageURL)
	if err != nil {
		return "", err
	}
	s, err := resp.Text()
	if err != nil {
		return "", err
	}

	reInput := regexp.MustCompile(
		`<input\s+[^>]*?name=['"]csrfmiddlewaretoken['"'][^>]*>`,
	)
	input := reInput.FindString(s)
	reValue := regexp.MustCompile(`value=['"]([^'"]+)['"]`)
	csrfToken := reValue.FindStringSubmatch(input)
	if len(csrfToken) < 2 {
		return "", err
	}
	return csrfToken[1], err
}

func login(req *request.Request) error {
	resp, err := req.Post(loginPageURL)
	defer resp.Body.Close() // **Don't forget close the response body**
	return err
}
