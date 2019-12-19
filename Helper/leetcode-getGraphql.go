package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsL4CodeSnippets []struct {
	Lang string
	Code string
}

type JsL3Detail struct {
	Content      string
	CodeSnippets JsL4CodeSnippets
}

type JsL2Question struct {
	Question JsL3Detail
}

type JsL1Data struct {
	Data JsL2Question
}

func getGraphql(p problem) (string, string) {
	// req := newReq()

	params := make(map[string]interface{})
	params["operationName"] = "questionData"
	params["query"] = "query questionData($titleSlug: String!) { question(titleSlug: $titleSlug) { content codeSnippets { lang code } } }"
	titleSlug := make(map[string]string)
	titleSlug["titleSlug"] = p.TitleSlug
	params["variables"] = titleSlug

	// Make this JSON
	postJson, _ := json.Marshal(params)

	// http.POST expects an io.Reader, which a byte buffer does
	postContent := bytes.NewBuffer(postJson)

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", postContent)
	if err != nil {
		log.Fatal("getGraphql: POST Error: " + err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("抓题失败 code: " + resp.Status)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("getGraphql: Read Error: " + err.Error())
	}

	//byte数组直接转成string，优化内存
	// str := (*string)(unsafe.Pointer(&respBytes))
	// fmt.Printf("%#v\n", *str)

	res := &JsL1Data{}
	json.Unmarshal(respBytes, &res)

	code := ""
	for _, qc := range res.Data.Question.CodeSnippets {
		if qc.Lang == "Go" {
			code = qc.Code
		}
	}
	return res.Data.Question.Content, code
}
