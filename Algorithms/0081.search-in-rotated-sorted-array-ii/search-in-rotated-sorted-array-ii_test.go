package problem0081

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
	nums   []int
	target int
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0081(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2},
				6,
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				6,
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				3,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[]int{},
				3,
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

		ast.Equal(a.one, search(p.nums, p.target), "输入:%v", p)
	}
}
