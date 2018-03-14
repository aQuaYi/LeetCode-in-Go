package problem0123

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
	prices []int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0123(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			ans{
				13,
			},
		},

		question{
			para{
				[]int{1, 2, 4, 2, 5, 7, 2, 3, 2, 3},
			},
			ans{
				8,
			},
		},

		question{
			para{
				[]int{},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{2, 1, 4},
			},
			ans{
				3,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
			ans{
				6,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 1, 2, 3, 1, 2, 3},
			},
			ans{
				4,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 1, 2, 3},
			},
			ans{
				4,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxProfit(p.prices), "输入:%v", p)
	}
}
