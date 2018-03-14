package problem0089

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
	one []int
}

func Test_Problem0089(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				4,
			},
			ans{
				[]int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8},
			},
		},

		question{
			para{
				3,
			},
			ans{
				[]int{0, 1, 3, 2, 6, 7, 5, 4},
			},
		},

		question{
			para{
				2,
			},
			ans{
				[]int{0, 1, 3, 2},
			},
		},

		question{
			para{
				1,
			},
			ans{
				[]int{0, 1},
			},
		},
		question{
			para{
				0,
			},
			ans{
				[]int{0},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, grayCode(p.n), "输入:%v", p)
	}
}
