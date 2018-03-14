package problem0414

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

func Test_Problem0414(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2, 2, 5, 3, 5}},
			ans{2},
		},

		question{
			para{[]int{1, 2, 3}},
			ans{1},
		},

		question{
			para{[]int{2, 2, 3, 1}},
			ans{1},
		},

		question{
			para{[]int{2, 3, 3}},
			ans{3},
		},

		question{
			para{[]int{3, 2, 1}},
			ans{1},
		},

		question{
			para{[]int{2, 2, 3}},
			ans{3},
		},

		question{
			para{[]int{1, 2}},
			ans{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, thirdMax(p.one), "输入:%v", p)
	}
}
