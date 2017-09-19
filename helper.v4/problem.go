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

	// 根据题号，获取题目信息
	for _, p := range ps {
		if p.ID == problemNum {
			pb = p
			isFound = true
			break
		}
	}

	if !isFound {
		log.Printf("没有发现 %d 题，存在以下可能：1.此题不存在；2.需要付费；3.不在已关注的种类中。", problemNum)
		return
	}

	build(pb)
}

func build(p problem) {
	if p.IsAccepted && GoKit.Exist(p.Dir) {
		log.Fatalf("第 %d 题已经accepted，请**删除**或**重命名**  %s 文件夹后，再尝试。", p.ID, p.Dir)
	}

	// 对于没有 accepted 的题目，直接删除重建
	if err := os.RemoveAll(p.Dir); err != nil {
		log.Fatalln("无法删除目录", p.Dir)
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	err := os.Mkdir(p.Dir, 0755)
	if err != nil {
		log.Fatalf("无法创建目录，%s ：%s", p.Dir, err)
	}

	log.Printf("开始创建 %d %s 的文件夹...\n", p.ID, p.Title)

	creatREADME(p)

	fc, fcName, para, ans := getFunction(p.link())
	creatGo(p, fc)
	creatGoTest(p, fcName, para, ans)

	log.Printf("%d.%s 的文件夹，创建完毕。\n", p.ID, p.Title)
}

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

## 题目
%s

## 解题思路

见程序注释
`

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

	return strings.TrimSpace(doc.Find("div.question-description").Text())
}

func getFunction(URL string) (fc, fcName, para, ansType string) {
	data := getRaw(URL)
	str := string(data)

	i := strings.Index(str, "codeDefinition:")
	j := i + strings.Index(str[i:], "enableTestMode:")
	str = str[i:j]

	i = strings.Index(str, "'Go', 'defaultCode': ") + 21
	j = i + strings.Index(str[i:], "},")
	str = str[i:j]

	i = strings.Index(str, "func")
	str = "'" + str[i:]

	// fmt.Println("getFunction: ", str)

	i = strings.Index(str, "'")
	j = 5 + strings.Index(str[5:], "'")
	fc = str[i+1 : j]

	k := 0
	i0 := strings.Index(fc, " ")
	i = strings.Index(fc, "(")
	fcName = fc[i0+1 : i]

	j = strings.Index(fc, ")")
	k = strings.Index(fc, "{")
	para = strings.Replace(fc[i+1:j], ",", "\n", -1)
	ansType = fc[j+1 : k]

	fc = fc[:k] + "{\n\n}"

	return
}

func creatGo(p problem, function string) {
	fileFormat := `package %s

%s
`
	content := fmt.Sprintf(fileFormat, p.packageName(), function)
	filename := fmt.Sprintf("%s/%s.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGoTest(p problem, fcName, para, ansType string) {
	testCasesFormat := `var tcs = []struct {
	%s
	ans %s
}{

	{
			
		,
	},

	// 可以有多个 testcase
}`

	testCases := fmt.Sprintf(testCasesFormat, para, ansType)

	testFuncFormat := `
func Test_%s(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%s~~\n", tc)
		ast.Equal(tc.ans, %s(%s), "输入:%s", tc)
	}
}`
	tcPara := getTcPara(para)
	testFunc := fmt.Sprintf(testFuncFormat, fcName, `%v`, fcName, tcPara)

	benchFuncFormat := `
func Benchmark_%s(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			%s(%s)
		}
	}
}`
	benchFunc := fmt.Sprintf(benchFuncFormat, fcName, fcName, tcPara)

	fileFormat := `package %s

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
%s

%s

%s
`

	content := fmt.Sprintf(fileFormat, p.packageName(), testCases, testFunc, benchFunc)

	filename := fmt.Sprintf("%s/%s_test.go", p.Dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// 把 函数的参数 变成 tc 的参数
func getTcPara(para string) string {
	// 把 para 按行切分
	paras := strings.Split(para, "\n")

	// 把单个参数按空格，切分成参数名和参数类型
	temp := make([][]string, len(paras))
	for i := range paras {
		temp[i] = strings.Split(strings.TrimSpace(paras[i]), ` `)
	}

	// 在参数名称前添加 "tc." 并组合在一起
	res := ""
	for i := 0; i < len(temp); i++ {
		res += ", tc." + temp[i][0]
	}

	return res[2:]
}

func (p problem) packageName() string {
	return fmt.Sprintf("Problem%04d", p.ID)
}
