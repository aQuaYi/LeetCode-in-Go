package problem0053

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

func Test_Problem0053(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-2}},
			ans{-2},
		},

		question{
			para{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans{6},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxSubArray(p.one), "输入:%v", p)
	}
}
