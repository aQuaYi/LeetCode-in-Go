package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

const (
	loginPageURL = "https://leetcode.com/accounts/login/"
)

// req 带有 cookie ，用来请求 leetcode 上的个人数据
var req *request.Request

// 登录 leetcode
func signin() {
	// 先导入配置
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Hi, %s. I'm %s\n", cfg.Login, VERSION)

	log.Println("正在登录中...")

	// 对 req 赋值
	req = request.NewRequest(new(http.Client))

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
		"login":               cfg.Login,
		"password":            cfg.Password,
	}
	if err = login(req); err != nil {
		log.Fatal(err)
	}

	log.Println("成功登录")
}

func getCategoryData(name string) *data {
	URL := url(name)

	raw := getRaw(URL)

	res := new(data)
	if err := json.Unmarshal(raw, res); err != nil {
		log.Fatal("无法把json转换成Category: " + err.Error())
	}

	return res
}

func url(s string) string {
	format := "https://leetcode.com/api/problems/%s/"
	return fmt.Sprintf(format, s)
}

func getRanking(username string) int {
	URL := fmt.Sprintf("https://leetcode.com/%s/", username)

	data := getRaw(URL)
	str := string(data)

	// fmt.Println(str)

	i := strings.Index(str, "ng-init")
	j := i + strings.Index(str[i:], "ng-cloak")
	str = str[i:j]

	i = strings.Index(str, "(")
	j = strings.Index(str, ")")
	str = str[i:j]

	//	fmt.Println("2\n", str)

	strs := strings.Split(str, ",")
	str = strs[6]

	//	fmt.Println("1\n", str)

	i = strings.Index(str, "'")
	j = 2 + strings.Index(str[2:], "'")

	//	fmt.Println("0\n", str)

	str = str[i+1 : j]

	r, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("无法把 %s 转换成数字Ranking", str)
	}

	return r
}

func getRaw(URL string) []byte {
	log.Printf("开始下载 %s 的数据", URL)

	resp, err := req.Get(URL)
	if err != nil {
		log.Fatal("getRaw: Get Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("getRaw: Read Error: " + err.Error())
	}
	return body
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

type config struct {
	Login, Password string
}
