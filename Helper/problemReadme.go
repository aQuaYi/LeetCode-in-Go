package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

## 题目

%s

## 解题思路

见程序注释
`

	questionDescription := strings.TrimSpace(getDescription(p.link()))

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.link(), questionDescription)

	content = replaceCharacters(content)

	filename := fmt.Sprintf("%s/README.md", p.Dir())

	write(filename, content)

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

func replaceCharacters(s string) string {
	changeMap := map[string]string{
		"&quot;":     "\"",
		"&lt;":       "<",
		"&gt;":       ">",
		"&ge;":       ">=",
		"&nbsp;":     "`",
		"&#39;":      "'",
		"&amp;":      "&",
		"   \n":      "\n",
		"  \n":       "\n",
		" \n":        "\n",
		"\n\n\n\n\n": "\n\n",
		"\n\n\n\n":   "\n\n",
		"\n\n\n":     "\n\n",
	}
	for old, new := range changeMap {
		s = strings.Replace(s, old, new, -1)
	}
	return s
}
