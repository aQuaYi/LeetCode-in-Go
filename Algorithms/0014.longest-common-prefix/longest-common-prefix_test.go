package problem0014

import (
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
	one string
}

func Test_Problem0014(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]string{"abcdd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		question{
			para{
				[]string{"abcdd", "abcde"},
			},
			ans{"abcd"},
		},
		question{
			para{
				[]string{},
			},
			ans{""},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestCommonPrefix(p.one), "输入:%v", p)
	}
}
