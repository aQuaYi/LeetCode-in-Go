package problem0096

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
	n int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0096(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				1,
			},
			ans{
				1,
			},
		},

		question{
			para{
				2,
			},
			ans{
				2,
			},
		},

		question{
			para{
				3,
			},
			ans{
				5,
			},
		},

		question{
			para{
				4,
			},
			ans{
				14,
			},
		},

		question{
			para{
				5,
			},
			ans{
				42,
			},
		},

		question{
			para{
				6,
			},
			ans{
				132,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, numTrees(p.n), "输入:%v", p)
	}
}
