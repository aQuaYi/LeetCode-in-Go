package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	fc, fcHead, para, ans := getFunction(p.link())

	creatREADME(p)
	creatGo(p, fc)
	creatGoTest(p, fcHead, para, ans)
}

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

## 题目
%s

## 解题思路


## 总结


`

	log.Printf("正在下载 %d.%s 的问题描述\n", p.ID, p.Title)

	questionDescription := getQuestionDescription(p.link())

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.link(), questionDescription)

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

func getFunction(URL string) (fc, fcHead, p, a string) {
	data := getRaw(URL)
	str := string(data)
	i := strings.Index(str, "codeDefinition:")
	j := i + strings.Index(str[i:], "enableTestMode:")
	str = str[i:j]

	i = strings.Index(str, "'func")
	j = i + strings.Index(str[i:], "},")
	str = str[i:j]

	i = strings.Index(str, "'")
	j = 5 + strings.Index(str[5:], "'")
	fc = str[i+1 : j]

	k := 0
	i0 := strings.Index(fc, " ")
	i = strings.Index(fc, "(")
	fcHead = fc[i0+1 : i]

	j = strings.Index(fc, ")")
	k = strings.Index(fc, "{")
	p = strings.Replace(fc[i+1:j], ",", "\n", -1)
	a = fc[j+1 : k]

	fc = fc[:k] + "{\n\n}"
	return
}

func creatGo(p problem, fc string) {
	fileFormat := `package %s

%s
`
	content := fmt.Sprintf(fileFormat, p.packageName(), fc)
	filename := fmt.Sprintf("%s/%s.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGoTest(p problem, fcHead, para, ans string) {
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
type para struct {
	%s
}

// ans 是答案
type ans struct {
	one %s
}

func Test_%s(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
					,
			},
			ans{
					,
			},
		},
	
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%s~~\n", p)

		ast.Equal(a.one, %s(p.one), "输入:%s", p)
	}
}
`
	content := fmt.Sprintf(fileFormat, p.packageName(), para, ans, p.packageName(), `%v`, fcHead, `%v`)
	filename := fmt.Sprintf("%s/%s_test.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func (p problem) packageName() string {
	return fmt.Sprintf("Problem%04d", p.ID)
}
