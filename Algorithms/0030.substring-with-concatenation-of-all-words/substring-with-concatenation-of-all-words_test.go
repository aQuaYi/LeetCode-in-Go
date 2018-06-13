package problem0030

import (
	"fmt"
	"sort"
	"testing"

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
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0030(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{"barfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{0, 9}},
		},

		question{
			para{"barbarfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{3, 12}},
		},

		question{
			para{"attoinattoin", []string{"at", "tto", "in"}},
			ans{[]int{}},
		},

		question{
			para{"attoinattoin", []string{"at", "to", "in"}},
			ans{[]int{0, 2, 4, 6}},
		},

		// leetcode的测试中，words中的单词长度大于0
		question{
			para{"", []string{""}},
			ans{[]int{}},
		},

		question{
			para{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}},
			ans{[]int{8}},
		},

		question{
			para{"barfoothefoobarmanattoinin", []string{"at", "to", "in", "in"}},
			ans{[]int{18}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		res := findSubstring(p.one, p.two)
		sort.Ints(res)
		ast.Equal(a.one, res, "输入:%v", p)
	}
}
