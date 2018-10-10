package problem0049

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
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]string
}

func Test_Problem0049(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans{[][]string{
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		res := groupAnagrams(p.one)
		for _, v := range res {
			sort.Strings(v)
			ast.Equal(a.one[len(v)-1], v, "输入:%v", p)
		}
	}
}
