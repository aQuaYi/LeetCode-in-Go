package problem0283

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
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0283(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 0, 1, 0, 3, 12}},
			ans{[]int{1, 1, 3, 12, 0, 0}},
		},
		question{
			para{[]int{0, 1, 0, 3, 12}},
			ans{[]int{1, 3, 12, 0, 0}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		moveZeroes(p.nums)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}
