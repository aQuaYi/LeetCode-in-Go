package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"syscall"

	"github.com/aQuaYi/GoKit"
)

func (p Problem) checkDir(CategoryDir string) string {
	pDir := fmt.Sprintf("%04d.%s", p.ID, p.TitleSlug)
	dir := fmt.Sprintf("./%s/%s", CategoryDir, pDir)

	if GoKit.Exist(dir) {
		return dir
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatalf("无法创建目录: %s", dir)
	}
	creatREADME(p, dir)
	creatGo(p, dir)
	creatGoTest(p, dir)

	return dir
}
func creatREADME(p Problem, dir string) {
	fileFormat := `# [%s](%s)
## 题目


## 解题思路


## 总结


`
	link := fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
	content := fmt.Sprintf(fileFormat, p.Title, link)
	filename := fmt.Sprintf("%s/README.md", dir)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGo(p Problem, dir string) {
	fileFormat := `package %s

`
	content := fmt.Sprintf(fileFormat, p.packageName())
	filename := fmt.Sprintf("%s/%s.go", dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGoTest(p Problem, dir string) {
	fileFormat := `package %s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	expected := ""
	actual := ""

	ast.Equal(expected, actual, "与预期不符")
}
`
	content := fmt.Sprintf(fileFormat, p.packageName())
	filename := fmt.Sprintf("%s/%s_test.go", dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// Problem 保存单个问题的信息
type Problem struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	PaidOnly   bool `json:"paid_only"`
	Difficulty `json:"difficulty"`
}

func (p Problem) packageName() string {
	return fmt.Sprintf("Problem%04d", p.ID)
}

// State 保存单个问题的解答状态
type State struct {
	Title     string `json:"question__title"`
	TitleSlug string `json:"question__title_slug"`
	IsNew     bool   `json:"is_new_question"`
	ID        int    `json:"question_id"`
	ACs       int    `json:"total_acs"`
	Submitted int    `json:"total_submitted"`
	PassRate  string
}

// Difficulty 问题的难度
type Difficulty struct {
	Level int `json:"level"`
}

// Solved 是已解决问题的缩写
type Solved struct {
	ID       int
	Title    string
	Dir      string
	Degree   string
	PassRate string
}

func (s Solved) String() string {
	res := fmt.Sprintf("|%d|", s.ID)
	res += fmt.Sprintf(`[%s](%s)|`, s.Title, s.Dir)
	res += fmt.Sprintf("%s|", s.Degree)
	res += fmt.Sprintf("%s|", s.PassRate)
	return res
}

var degrees = map[int]string{
	1: ` ☆ `,
	2: ` ☆  ☆ `,
	3: ` ☆  ☆  ☆ `,
}

func makeSolved(p Problem, dir string) Solved {
	return Solved{
		ID:       p.ID,
		Title:    p.Title,
		Dir:      fmt.Sprintf("%s", dir),
		Degree:   degrees[p.Difficulty.Level],
		PassRate: p.PassRate,
	}
}

// Solveds 保存所有已经解答的问题
type Solveds []Solved

func (ss Solveds) Len() int {
	return len(ss)
}

func (ss Solveds) Less(i, j int) bool {
	return ss[i].ID < ss[j].ID
}

func (ss Solveds) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss Solveds) String() string {
	sort.Sort(ss)
	res := "|题号|题目|难度|通过率|\n"
	res += "|:-:| :-- | :-: | :-: |\n"
	for _, s := range ss {
		res += fmt.Sprintln(s)
	}
	return res
}
