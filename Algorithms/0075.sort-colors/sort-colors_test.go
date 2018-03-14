package problem0075

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
	one []int
}

func Test_Problem0075(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{0, 2, 0, 2, 0, 2},
			},
			ans{
				[]int{0, 0, 0, 2, 2, 2},
			},
		},

		question{
			para{
				[]int{1, 2, 1, 2, 1, 2},
			},
			ans{
				[]int{1, 1, 1, 2, 2, 2},
			},
		},

		question{
			para{
				[]int{0, 1, 2, 0, 1, 2, 0, 1, 2},
			},
			ans{
				[]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
			},
		},

		question{
			para{
				[]int{2, 1},
			},
			ans{
				[]int{1, 2},
			},
		},

		question{
			para{
				[]int{},
			},
			ans{
				[]int{},
			},
		},

		question{
			para{
				[]int{2, 2, 2, 2, 2, 2},
			},
			ans{
				[]int{2, 2, 2, 2, 2, 2},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		sortColors(p.nums)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}
