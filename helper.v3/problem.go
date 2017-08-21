package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/aQuaYi/GoKit"

	"github.com/PuerkitoBio/goquery"
)

func makeProblemDir(ps problems, problemNum int) {
	var pb problem
	var isFound bool
	for _, p := range ps {
		if p.ID == problemNum {
			pb = p
			isFound = true
			break
		}
	}

	if !isFound {
		log.Printf("没有发现 %d 题，此题不存在，或者需要付费，或者不在已关注的种类中。", problemNum)
		return
	}

	if pb.IsAccepted && GoKit.Exist(pb.Dir) {
		log.Printf("第 %d 题已经accepted，请删除 %s 文件夹后，再尝试。", pb.ID, pb.Dir)
		return
	}

	pb.build()
}

func (p problem) build() {
	if err := os.RemoveAll(p.Dir); err != nil {
		log.Fatalln("无法删除目录", p.Dir)
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	err := os.Mkdir(p.Dir, 0755)
	if err != nil {
		log.Fatalf("无法创建目录: %s", p.Dir)
	}

	creatREADME(p)
	creatGo(p)
	creatGoTest(p)
}

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

## 题目
%s

## 解题思路


## 总结


`
	link := fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)

	log.Printf("正在下载 %d.%s 的问题描述\n", p.ID, p.Title)

	questionDescription := getQuestionDescription(link)

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, link, questionDescription)

	filename := fmt.Sprintf("%s/README.md", p.Dir)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getQuestionDescription(URL string) string {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("div.question-description").Text()
}

func creatGo(p problem) {
	fileFormat := `package %s

`
	content := fmt.Sprintf(fileFormat, p.packageName())
	filename := fmt.Sprintf("%s/%s.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGoTest(p problem) {
	fileFormat := `package %s

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one string
}

func Test_%s(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{""},
			ans{""},
		},
	
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%s~~\n", p)

		ast.Equal(a.one, (p.one), "输入:%s", p)
	}
}
`
	content := fmt.Sprintf(fileFormat, p.packageName(), p.packageName(), `%v`, `%v`)
	filename := fmt.Sprintf("%s/%s_test.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func (p problem) packageName() string {
	return fmt.Sprintf("Problem%04d", p.ID)
}
