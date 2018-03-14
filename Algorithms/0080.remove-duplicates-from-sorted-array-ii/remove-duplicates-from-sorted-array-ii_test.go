package problem0080

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
	one   int
	first []int
}

func Test_Problem0080(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 1, 1, 2, 2, 3},
			},
			ans{
				5,
				[]int{1, 1, 2, 2, 3},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
			},
			ans{
				5,
				[]int{1, 2, 3, 4, 5},
			},
		},

		question{
			para{
				[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			ans{
				2,
				[]int{1, 1},
			},
		},

		question{
			para{
				[]int{1, 1},
			},
			ans{
				2,
				[]int{1, 1},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		end := removeDuplicates(p.nums)
		ast.Equal(a.one, end, "输入:%v", p)
		ast.Equal(a.first, p.nums[:end], "输入:%v", p)
	}
}
