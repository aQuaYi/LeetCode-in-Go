package problem0152

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
	nums []int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0152(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]int{-5, -4, -3, -2, -1, -2, -3, -4, -5},
			},
			ans{
				2880,
			},
		},

		question{
			para{
				[]int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			},
			ans{
				14400,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, -5, 6, 7, 8, 9, 10},
			},
			ans{
				30240,
			},
		},

		question{
			para{
				[]int{-2, 0, -2, 0, -2, 0, -2, 0, -2, 0, -2},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{2, -5, -2, -4, 3},
			},
			ans{
				24,
			},
		},

		question{
			para{
				[]int{3, -5, -3, -4, 3},
			},
			ans{
				45,
			},
		},

		question{
			para{
				[]int{3, -1, 4},
			},
			ans{
				4,
			},
		},

		question{
			para{
				[]int{0, 0, 2},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[]int{2, 3, -2, 4},
			},
			ans{
				6,
			},
		},

		question{
			para{
				[]int{1, 2, 3, -4, -5, 6},
			},
			ans{
				720,
			},
		},

		question{
			para{
				[]int{-2, 0, 2},
			},
			ans{
				2,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxProduct(p.nums), "输入:%v", p)
	}
}
