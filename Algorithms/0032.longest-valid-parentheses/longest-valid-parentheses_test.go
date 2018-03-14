package problem0032

import (
	"fmt"
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
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0032(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{")(()()))((((())))("},
			ans{8},
		},

		question{
			para{"()(()"},
			ans{2},
		},

		question{
			para{"(()())"},
			ans{6},
		},

		question{
			para{"(()"},
			ans{2},
		},

		question{
			para{")()())"},
			ans{4},
		},

		question{
			para{"((()))"},
			ans{6},
		},

		question{
			para{"((())))(((())))"},
			ans{8},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, longestValidParentheses(p.one), "输入:%v", p)
	}
}
