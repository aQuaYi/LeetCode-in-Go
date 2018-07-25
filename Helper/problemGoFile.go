package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func parseFunction(fc string) (fcName, para, ansType, nfc string) {
	log.Println("准备分解函数:", fc)

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fcName = "myFunc"
			para = "p int"
			ansType = "int"
			nfc = "func myFunc(p int) int {\n\n}"
		}
	}()

	funcIndex := strings.Index(fc, "func ")
	a := funcIndex + strings.Index(fc[funcIndex:], " ")
	b := funcIndex + strings.Index(fc[funcIndex:], "(")
	c := funcIndex + strings.Index(fc[funcIndex:], ")")
	d := funcIndex + strings.Index(fc[funcIndex:], "{")

	fcName = fc[a+1 : b]
	para = fc[b+1 : c]
	ansType = strings.TrimSpace(fc[c+1 : d])
	nfc = fmt.Sprintf("func %s(%s) %s {\n\n}", fcName, para, ansType)

	return
}

func getFunction(url string) string {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run task list
	var function string

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Println("chromedp.New 出错：", err)
	}

	err = c.Run(ctxt, makeTasks(url, &function))
	if err != nil {
		log.Println("c.Run 出错：", err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Println("c.Shutdown 出错：", err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Println("c.Wait 出错：", err)
	}

	log.Println("抓取到函数：", function)

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
