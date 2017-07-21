package main

import (
	"io/ioutil"
	"log"
	"syscall"

	"os"

	"fmt"

	"sort"

	"strings"

	"github.com/aQuaYi/GoKit"
)

func (c *Category) run() Solveds {
	checkUser(c.User)
	checkDir(c.Name)
	c.compute()
	res := c.analysis()
	fmt.Println(res)
	return res
}

// 根据已有的内容，填上空缺字段
func (c *Category) compute() {
	for _, p := range c.Problems {
		// 统计各个难度的题目数量
		switch p.Difficulty.Level {
		case 1:
			c.Easy++
		case 2:
			c.Medium++
		case 3:
			c.Hard++
		default:
			log.Fatalln("出现了第4种难度")
		}

		p.PassRate = fmt.Sprintf("%d%%", 100*p.ACs/p.Submitted)
		fmt.Println(p.PassRate, p)
	}

	// 统计AC的总数
	c.AC = c.ACEasy + c.ACMedium + c.ACHard

	if c.Total != c.Easy+c.Medium+c.Hard {
		log.Fatalf("%s分类下的各难度题目数量之和，不等于Total", c.Name)
	}
}

func (c *Category) analysis() Solveds {
	res := Solveds{}

	for _, p := range c.Problems {
		pdir := p.checkDir(c.Name)
		// 检查处理已经AC的题目
		if p.Status == "ac" {
			res = append(res, makeSolved(*p, pdir))
		}
	}

	return res
}

func (p Problem) checkDir(CategoryDir string) string {
	pDir := fmt.Sprintf("%d.%s", p.ID, p.TitleSlug)
	dir := fmt.Sprintf("./%s/%s", CategoryDir, pDir)
	fmt.Println(dir)

	if GoKit.Exist(dir) {
		fmt.Println("存在", dir)
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
	packageName := strings.Replace(p.Title, " ", "", -1)
	fileFormat := `package %s

`
	content := fmt.Sprintf(fileFormat, packageName)
	filename := fmt.Sprintf("%s/%s.go", dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func creatGoTest(p Problem, dir string) {
	packageName := strings.Replace(p.Title, " ", "", -1)
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
	content := fmt.Sprintf(fileFormat, packageName)
	filename := fmt.Sprintf("%s/%s_test.go", dir, p.TitleSlug)

	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

type Category struct {
	Name     string `json:"category_slug"`
	User     string `json:"user_name"`
	ACEasy   int    `json:"ac_easy"`
	ACMedium int    `json:"ac_medium"`
	ACHard   int    `json:"ac_hard"`
	AC       int    `json:"num_solved"`
	Easy     int
	Medium   int
	Hard     int
	Total    int        `json:"num_total"`
	Problems []*Problem `json:"stat_status_pairs"`
}

type Problem struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	PaidOnly   bool `json:"paid_only"`
	Difficulty `json:"difficulty"`
}

type State struct {
	Title     string `json:"question__title"`
	TitleSlug string `json:"question__title_slug"`
	IsNew     bool   `json:"is_new_question"`
	ID        int    `json:"question_id"`
	ACs       int    `json:"total_acs"`
	Submitted int    `json:"total_submitted"`
	PassRate  string
}

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
	1: `:star:`,
	2: `:star::star:`,
	3: `:star::star::star:`,
}

func makeSolved(p Problem, dir string) Solved {
	fmt.Println("in makeSolved", p.PassRate)
	return Solved{
		ID:       p.ID,
		Title:    p.Title,
		Dir:      fmt.Sprintf("%s", dir),
		Degree:   degrees[p.Difficulty.Level],
		PassRate: p.PassRate,
	}
}

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
	res += "|---| ----- | -------- | ---------- |\n"
	for _, s := range ss {
		res += fmt.Sprintln(s) + "\n"
	}
	return res
}

func checkUser(u string) {
	if u != USER {
		log.Fatalln("下载的不是本人的数据，请按照helper的说明文档，重新获取leetcode.cookie")
	}
}

func checkDir(dir string) {
	if GoKit.Exist(dir) {
		return
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatalf("无法创建文件夹%s", dir)
	}

	log.Printf("已经创建文件夹%s", dir)
}
