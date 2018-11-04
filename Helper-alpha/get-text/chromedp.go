// Command eval is a chromedp example demonstrating how to evaluate javascript
// and retrieve the result.
package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	url := "https://leetcode.com/problems/valid-permutations-for-di-sequence/"
	res := getDescription(url)
	log.Printf("main:%s ", res)
}

func getDescription(url string) string {
	log.Printf("准备访问 %s", url)

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
	var res string
	err = c.Run(ctxt, text(url, &res))
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

	res = deleteLineNumber(res)
	res = deleteMoreBlank(res)

	res = convert(res)

	return res
}

func text(url string, res *string) chromedp.Tasks {
	codeSel := "#main-container > div > div > div.editor-wrapper__1v3H > div > div.content__1YQ2 > div > div.container__YPDh > div > div > div.CodeMirror-scroll > div.CodeMirror-sizer > div > div > div > div.CodeMirror-code"
	return chromedp.Tasks{
		chromedp.Navigate(url),
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

func deleteMoreBlank(f string) string {
	for strings.Contains(f, "  ") {
		f = strings.Replace(f, "  ", " ", 1)
	}
	return f
}

func convert(f string) string {
	publicIndex := strings.Index(f, "public")

	f = f[publicIndex:]

	fs := strings.Split(f, " ")

	returnType := fs[2]
	fs3 := fs[3]
	bIndex := strings.Index(fs3, "(")
	funcName := fs3[:bIndex]

	paras := "ps"

	for i := range fs {
		fmt.Printf("%d: %s\n", i, fs[i])
	}

	format := "func %s(%s) %s { }"

	return fmt.Sprintf(format, funcName, paras, returnType)
}
