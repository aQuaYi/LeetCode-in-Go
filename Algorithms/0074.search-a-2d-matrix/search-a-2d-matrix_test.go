package problem0074

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
	matrix [][]int
	target int
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0074(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[][]int{},
				3,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[][]int{[]int{}},
				3,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				3,
			},
			ans{
				true,
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				0,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				4,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 50},
				},
				36,
			},
			ans{
				false,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, searchMatrix(p.matrix, p.target), "输入:%v", p)
	}
}
