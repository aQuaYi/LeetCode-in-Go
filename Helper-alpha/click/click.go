package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	var function string

	// run task list
	err = c.Run(ctxt, click(&function))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	function = deleteLineNumber(function)

	fmt.Println(function)
}

func click(res *string) chromedp.Tasks {
	btn := `#main-container > div > div > div.editor-wrapper__1v3H > div > div.content__1YQ2 > div > div.container__2zYY > div.select__2iyW.select-container__2w2b > div > div`
	codeSel := "#main-container > div > div > div.editor-wrapper__1v3H > div > div.content__1YQ2 > div > div.container__YPDh > div > div > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code"
	return chromedp.Tasks{
		chromedp.Navigate(`https://leetcode.com/problems/increasing-order-search-tree/`),

		// TODO: How can I choose the "Go" selection

		chromedp.Click(btn, chromedp.ByQuery),

		chromedp.Sleep(1000 * time.Second),

		chromedp.Text(codeSel, res),
	}
}

func deleteLineNumber(f string) string {
	i := 1
	for {
		si := strconv.Itoa(i)
		if !strings.Contains(f, si) {
			return f
		}
		f = strings.Replace(f, si, "", 1)
		i++
	}
}
