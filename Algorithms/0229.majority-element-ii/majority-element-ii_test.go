package problem0229

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

func Test_Problem0229(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3, 1, 1, 3, 3, 2},
			},
			ans{
				[]int{1, 3},
			},
		},

		question{
			para{
				[]int{0, 3, 4, 0},
			},
			ans{
				[]int{0},
			},
		},

		question{
			para{
				[]int{1},
			},
			ans{
				[]int{1},
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
