package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/aQuaYi/GoKit"
)

func buildProblemDir(problemNum int) {
	log.Printf("~~ 开始生成第 %d 题的文件夹 ~~\n", problemNum)

	// 需要创建答题文件夹
	lc := newLeetCode()
	//
	makeProblemDir(lc.Problems, problemNum)
	//
	log.Printf("~~ 第 %d 题的文件夹，已经生成 ~~\n", problemNum)
}

func makeProblemDir(ps problems, problemNum int) {
	var pb problem
	var isFound bool

	// 根据题号，获取题目信息
	for _, p := range ps {
		if p.ID == problemNum {
			if p.HasNoGoOption {
				log.Fatalln(`此题被标记为"不能使用 Go 语言解答"。请核查后，修改 unavailable.json 中的记录`)
			}
			pb = p
			isFound = true
			break
		}
	}

	if !isFound {
		log.Printf("没有发现第 %d 题，存在以下可能：1.此题不存在；2.此题需要付费。", problemNum)
		return
	}

	// 创建目录
	build(pb)
}

func build(p problem) {
	if GoKit.Exist(p.Dir()) {
		log.Fatalf("第 %d 题的文件夹已经存在，请**移除**  %s 文件夹后，再尝试。", p.ID, p.Dir())
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	// 创建目录
	err := os.Mkdir(p.Dir(), 0755)
	if err != nil {
		log.Fatalf("无法创建目录，%s ：%s", p.Dir(), err)
	}

	log.Printf("开始创建 %d %s 的文件夹...\n", p.ID, p.Title)

	creatREADME(p)

	fc := getFunction(p.link())

	fcName, para, ans := parseFunction(fc)

	creatGo(p, fc, ans)

	creatGoTest(p, fcName, para, ans)

	// 利用 chrome 打开题目 submissions 页面
	go func() {
		cmd := exec.Command("google-chrome", "https://leetcode.com/submissions/")
		_, err := cmd.Output()
		if err != nil {
			panic(err.Error())
		}
	}()

	log.Println("等待 10 秒，打开题目页面")
	time.Sleep(10 * time.Second)

	// 利用 chrome 打开题目页面
	go func() {
		cmd := exec.Command("google-chrome", p.link())
		_, err := cmd.Output()
		if err != nil {
			panic(err.Error())
		}
	}()

	log.Println("正在打开题目页面")
	time.Sleep(2 * time.Second)

	log.Printf("%d.%s 的文件夹，创建完毕。\n", p.ID, p.Title)
}

var typeMap = map[string]string{
	"int":     "0",
	"float64": "0",
	"string":  "\"\"",
	"bool":    "false",
}

func creatGo(p problem, function, ansType string) {
	fileFormat := `package %s

%s
`
	content := fmt.Sprintf(fileFormat, p.packageName(), function)

	returns := "\treturn nil\n}"
	if v, ok := typeMap[ansType]; ok {
		returns = fmt.Sprintf("\treturn %s\n}", v)
	}

	content = strings.Replace(content, "}", returns, -1)

	filename := fmt.Sprintf("%s/%s.go", p.Dir(), p.TitleSlug)

	write(filename, content)
}

func creatGoTest(p problem, fcName, para, ansType string) {
	testCasesFormat := `var tcs = []struct {
	%s
	ans %s
}{


	
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
	testFunc := fmt.Sprintf(testFuncFormat, fcName, `%v`, fcName, tcPara, `%v`)

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

	filename := fmt.Sprintf("%s/%s_test.go", p.Dir(), p.TitleSlug)

	write(filename, content)

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
	return fmt.Sprintf("problem%04d", p.ID)
}
