package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/aQuaYi/GoKit"
)

func buildProblemDir(problemNum int) {
	log.Printf("~~ 开始生成第 %d 题的文件夹 ~~\n", problemNum)

	// 获取 LeetCode 的记录文件
	lc := newLeetCode()

	// 检查 problemNum 的合法性
	if problemNum >= len(lc.Problems) {
		log.Panicf("%d 超出题目范围，请核查题号。", problemNum)
	}
	if lc.Problems[problemNum].ID == 0 {
		log.Panicf("%d 号题不存，请核查题号。", problemNum)
	}
	if lc.Problems[problemNum].IsPaid {
		log.Panicf("%d 号题需要付费。如果已经订阅，请注释掉本代码。", problemNum)
	}
	if lc.Problems[problemNum].HasNoGoOption {
		log.Panicf("%d 号题，没有提供 Go 解答选项。请核查后，修改 unavailable.json 中的记录。", problemNum)
	}

	// 需要创建答题文件夹
	build(lc.Problems[problemNum])

	log.Printf("~~ 第 %d 题的文件夹，已经生成 ~~\n", problemNum)
}

func build(p problem) {
	if GoKit.Exist(p.Dir()) {
		log.Panicf("第 %d 题的文件夹已经存在，请 **移除** %s 文件夹后，再尝试。", p.ID, p.Dir())
	}

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Println(err)
			log.Println("清理不必要的文件")
			os.RemoveAll(p.Dir())
		}
	}()

    // windows用户注释这两行
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	// 创建目录
	err := os.Mkdir(p.Dir(), 0755)
	if err != nil {
		log.Panicf("无法创建目录，%s ：%s", p.Dir(), err)
	}

	log.Printf("开始创建 %d %s 的文件夹...\n", p.ID, p.Title)

	content, fc := getGraphql(p)
	if fc == "" {
		log.Panicf("查无Go语言写法")
	}

	// 利用 chrome 打开题目页面
	// go func() {
	// 	cmd := exec.Command("google-chrome", p.link())
	// 	_, err = cmd.Output()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }()

	// fc := getFunction(p.link())

	fcName, para, ans, _ := parseFunction(fc)

	creatGo(p, fc, ans)

	creatGoTest(p, fcName, para, ans)

	creatREADME(p, content)

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
%s
`

	treeNodeDefine := ""
	if strings.Contains(function, "*TreeNode") {
		treeNodeDefine = `
import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

`
	}

	content := fmt.Sprintf(fileFormat, p.packageName(), treeNodeDefine, function)

	if v, ok := typeMap[ansType]; ok {
		content = strings.Replace(content, "nil", v, 1)
	}

	filename := fmt.Sprintf("%s/%s.go", p.Dir(), p.TitleSlug)

	write(filename, content)

	vscodeOpen(filename)
}

func creatGoTest(p problem, fcName, para, ansType string) {
	testCasesFormat := `var tcs = []struct {
	%s
	ans %s
}{



	// 可以有多个 testcase
}`

	para = strings.Replace(para, ",", "\n", -1)

	testCases := fmt.Sprintf(testCasesFormat, para, ansType)

	testFuncFormat := `
func Test_%s(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, %s(%s), "输入:%s", tc)
	}
}`
	tcPara := getTcPara(para)
	testFunc := fmt.Sprintf(testFuncFormat, fcName, fcName, tcPara, `%v`)

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

	vscodeOpen(filename)

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
