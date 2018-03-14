package problem0055

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
	one bool
}

func Test_Problem0055(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{2, 3, 1, 1, 4},
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{3, 2, 1, 0, 4},
			},
			ans{
				false,
			},
		},

		question{
			para{
				[]int{3, 2, 2, 0, 4},
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{3, 2, 2, 0, 0},
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

		ast.Equal(a.one, canJump(p.nums), "输入:%v", p)
	}
}
