package problem0643

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
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one float64
}

func Test_Problem0643(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-1}, 1},
			ans{-1},
		},

		question{
			para{[]int{1, 12, -5, -6, 50, 3}, 4},
			ans{12.75},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findMaxAverage(p.nums, p.k), "输入:%v", p)
	}
}
