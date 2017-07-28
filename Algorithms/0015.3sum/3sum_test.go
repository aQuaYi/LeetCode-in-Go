package Problem0015

import (
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
	one [][]int
}

func Test_Problem0015(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{-1, 0, 1, 2, -1, -4}},
			ans{[][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
	}
}
