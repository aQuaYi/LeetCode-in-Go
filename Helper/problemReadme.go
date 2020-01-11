package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/TruthHun/html2md"
)

func creatREADME(p problem, s string) {
	fileFormat := `# [%d. %s](%s)

%s
`

	questionDescription := strings.TrimSpace(getDescription(p.link()))

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.link(), questionDescription) + s + "\n\n## 解题思路\n\n## 可能的變化"

	content = replaceCharacters(content)

	content = html2md.Convert(content)

	filename := fmt.Sprintf("%s/README.md", p.Dir())

	write(filename, content)

	vscodeOpen(filename)

}

func replaceCharacters(s string) string {
	changeMap := map[string]string{
		"&amp;lt;":   "<",
		"&amp;quot;": "\"",
		"&amp;nbsp;": " ",
		"&amp;#39;":  "`",
		"&quot;":     "\"",
		"&lt;":       "<",
		"&gt;":       ">",
		"&ge;":       ">=",
		"&nbsp;":     " ",
		"&amp;":      "&",
		"&#39;":      "'",
		"   \n":      "\n",
		"  \n":       "\n",
		" \n":        "\n",
		"\n\n\n\n\n": "\n\n",
		"\n\n\n\n":   "\n\n",
		"\n\n\n":     "\n\n",
	}

	olds := make([]string, 0, len(changeMap))
	for old := range changeMap {
		olds = append(olds, old)
	}

	sort.Strings(olds)

	news := make([]string, 0, len(olds))
	for _, old := range olds {
		news = append(news, changeMap[old])
	}

	for i := len(olds) - 1; 0 <= i; i-- {
		// 先替换长的，再替换短的
		old, new := olds[i], news[i]
		s = strings.Replace(s, old, new, -1)
	}

	return s

}

func getDescription(url string) string {
	return ""
}

// func getDescription(url string) string {
// 	var err error

// 	// create context
// 	ctxt, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	var options chromedp.Option
// 	options = chromedp.WithRunnerOptions(
// 		runner.Flag("headless", true),
// 		runner.Flag("no-sandbox", true),
// 		runner.Flag("disable-gpu", true),
// 	)

// 	log.Println("chromedp timeout:", chromedp.DefaultNewTargetTimeout)

// 	// create chrome instance
// 	c, err := chromedp.New(ctxt, options)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// run task list
// 	var res string
// 	err = c.Run(ctxt, text(url, &res))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// shutdown chrome
// 	err = c.Shutdown(ctxt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// wait for chrome to finish
// 	err = c.Wait()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Desc:", res)

// 	return res
// }

// func text(url string, res *string) chromedp.Tasks {
// 	sel := `div.content__eAC7`
// 	return chromedp.Tasks{
// 		chromedp.Sleep(time.Second * 3),
// 		chromedp.Navigate(url),
// 		chromedp.Text(sel, res, chromedp.NodeVisible, chromedp.BySearch),
// 	}
// }
