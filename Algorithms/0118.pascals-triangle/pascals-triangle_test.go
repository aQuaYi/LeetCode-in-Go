package problem0118

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
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0118(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{0},
			ans{[][]int{}},
		},

		question{
			para{1},
			ans{[][]int{
				[]int{1},
			}},
		},

		question{
			para{5},
			ans{[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, generate(p.one), "输入:%v", p)
	}
}
