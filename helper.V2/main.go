package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

const (
	loginPageURL = "https://leetcode.com/accounts/login/"
)

func algorithm(req *request.Request) string {
	resp, err := req.Get("https://leetcode.com/api/problems/Algorithms/")
	if err != nil {
		return "Get Error: " + err.Error()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Read Error: " + err.Error()
	}

	type Category struct {
		Total int    `json:"num_total"`
		User  string `json:"user_name"`
	}

	c := Category{}
	if err := json.Unmarshal(body, &c); err != nil {
		return "Unmarshal Error: " + err.Error()
	}

	return fmt.Sprintf("Your username is %s. Total number is %d.", c.User, c.Total)
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
	body, _ := ioutil.ReadAll(resp.Body)
	fr, _ := os.Create("login.html")
	fr.Write(body)
	return err
}

type config struct {
	Login, Password string
}

func main() {
	cfg := new(config)
	if _, err := toml.DecodeFile("leetcode.toml", cfg); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(*cfg)

	c := new(http.Client)
	req := request.NewRequest(c)
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

	log.Println("algorithm: ", algorithm(req))
}
