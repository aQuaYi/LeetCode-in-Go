package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

## 题目

%s
## 解题思路

见程序注释
`

	questionDescription := getDescription(p.link())

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.link(), questionDescription)

	filename := fmt.Sprintf("%s/README.md", p.Dir)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getDescription(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var desc string

	doc.Find("meta[name=description]").Each(func(i int, selection *goquery.Selection) {
		desc, _ = selection.Attr("content")
	})

	return desc
}
