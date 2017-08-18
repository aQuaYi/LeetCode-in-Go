package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func showRanking(username string) {
	URL := fmt.Sprintf("https://leetcode.com/%s", username)
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	ranking, _ := doc.Find("div.content-wrapper").Find("div.response-container container ng-scope").Attr("ng-app")

	fmt.Printf("%s's ranking is %v", username, ranking)
}

func main() {
	showRanking("aQuaYi")
}
