package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func parseFunction(fc string) (fcName, para, ansType string) {
	funcIndex := strings.Index(fc, "func ")
	a := funcIndex + strings.Index(fc[funcIndex:], " ")
	b := funcIndex + strings.Index(fc[funcIndex:], "(")
	c := funcIndex + strings.Index(fc[funcIndex:], ")")
	d := funcIndex + strings.Index(fc[funcIndex:], "{")

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
		log.Fatal("chromedp.New 出错：", err)
	}

	// run task list
	var function string
	err = c.Run(ctxt, makeTasks(url, &function))
	if err != nil {
		log.Fatal("c.Run 出错：", err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal("c.Shutdown 出错：", err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal("c.Wait 出错：", err)
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
