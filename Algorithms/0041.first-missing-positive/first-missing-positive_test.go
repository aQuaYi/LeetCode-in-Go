package problem0041

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
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0041(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{10, -1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, -3}},
			ans{6},
		},

		question{
			para{[]int{10, -1, 8, 6, 7, 3, -2, 5, 4, 2, 1, -3}},
			ans{9},
		},

		question{
			para{[]int{1}},
			ans{2},
		},

		question{
			para{[]int{0, 2, 2, 1, 1}},
			ans{3},
		},

		question{
			para{[]int{}},
			ans{1},
		},

		question{
			para{[]int{1, 2, 0}},
			ans{3},
		},

		question{
			para{[]int{3, 4, -1, 1}},
			ans{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, firstMissingPositive(p.one), "输入:%v", p)
	}
}
