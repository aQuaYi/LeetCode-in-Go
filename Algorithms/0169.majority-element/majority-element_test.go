package problem0169

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

func Test_Problem0169(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 1, 2, 2, 2, 1, 1, 2, 2},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[]int{1},
			},
			ans{
				1,
			},
		},

		question{
			para{
				[]int{1, 1, 2, 2, 2},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[]int{1, 2, 2},
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

		ast.Equal(a.one, majorityElement(p.nums), "输入:%v", p)
	}
}
