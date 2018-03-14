package problem0045

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

func Test_Problem0045(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{2, 3, 1, 4, 1, 1, 1, 4, 1, 5, 1, 1, 1, 1}},
			ans{5},
		},

		question{
			para{[]int{2, 3, 1, 1, 1, 1, 1, 4, 1, 5, 1, 1, 1, 1}},
			ans{7},
		},

		question{
			para{[]int{2, 3, 1, 4, 1, 5, 1, 1, 1, 1}},
			ans{4},
		},

		question{
			para{[]int{2, 3, 1, 1, 4}},
			ans{2},
		},

		question{
			para{[]int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}},
			ans{13},
		},

		question{
			para{[]int{0}},
			ans{0},
		},

		question{
			para{[]int{}},
			ans{0},
		},

		question{
			para{[]int{5, 1, 1, 1, 1, 1, 1}},
			ans{2},
		},

		question{
			para{[]int{1, 1, 1, 1, 1, 1}},
			ans{5},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, jump(p.one), "输入:%v", p)
	}
}
