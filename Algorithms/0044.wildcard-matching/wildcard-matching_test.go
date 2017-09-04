package Problem0044

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
type para struct {
	s string
	p string
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0044(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"aa",
				"a",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"aa",
				"aa",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"aaa",
				"aa",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"aa",
				"*",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"aa",
				"a*",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"ab",
				"?*",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"aab",
				"c*a*b",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"ab",
				"*",
			},
			ans{
				true,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isMatch(p.s, p.p), "输入:%v", p)
	}
}
