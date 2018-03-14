package problem0189

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
	k    int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0189(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6}, 2,
			},
			ans{
				[]int{5, 6, 1, 2, 3, 4},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 14,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 7,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			},
			ans{
				[]int{5, 6, 7, 1, 2, 3, 4},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		rotate(p.nums, p.k)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}
