package problem0087

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
	s1 string
	s2 string
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0087(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"abcd",
				"bdac",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"abcd",
				"hijk",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"abc",
				"abc",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"eat",
				"ate",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"a",
				"b",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"great",
				"rgeat",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"great",
				"rgtae",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"eat",
				"aet",
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

		ast.Equal(a.one, isScramble(p.s1, p.s2), "输入:%v", p)
	}
}
