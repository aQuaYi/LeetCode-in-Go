package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://leetcode.com/aQuaYi/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc)
	// Find the review items
	fmt.Println(doc.Find("panel-body").FindSelection(sel *goquery.Selection)("h4"))
}

func main() {
	ExampleScrape()
}
