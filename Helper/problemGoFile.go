package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func parseFunction(fc string) (fcName, para, ansType string) {
	a := strings.Index(fc, " ")
	b := strings.Index(fc, "(")
	c := strings.Index(fc, ")")
	d := strings.Index(fc, "{")

	fcName = fc[a+1 : b]
	para = strings.Replace(fc[b+1:c], ",", "\n", -1)
	ansType = strings.TrimSpace(fc[c+1 : d])

	return
}

func getFunction(url string) string {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var function string
	err = c.Run(ctxt, makeTasks(url, &function))
	if err != nil {
		log.Println("Run error")
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Println("Shutdown error")
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Println("Wait error")
		log.Fatal(err)
	}

	return function
}

func makeTasks(url string, function *string) chromedp.Tasks {
	textarea := `//textarea`
	btn := `#question-detail-app > div > div:nth-child(3) > div > div > div.row.control-btn-bar > div > div > div > div > span.Select-arrow-zone`
	goSel := `#react-select-2--option-9`
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Click(btn, chromedp.ByID),
		chromedp.Click(goSel, chromedp.ByID),
		chromedp.Text(textarea, function),
	}
}
