package problem0066

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
	one []int
}

func Test_Problem0066(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2, 3}},
			ans{[]int{1, 2, 4}},
		},

		question{
			para{[]int{0}},
			ans{[]int{1}},
		},

		question{
			para{[]int{9, 9}},
			ans{[]int{1, 0, 0}},
		},

		question{
			para{[]int{}},
			ans{[]int{1}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, plusOne(p.one), "输入:%v", p)
	}
}
