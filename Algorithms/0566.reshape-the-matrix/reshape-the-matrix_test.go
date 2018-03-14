package problem0566

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
	one  [][]int
	r, c int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0566(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			}, 2, 4},
			ans{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			}, 1, 4},
			ans{[][]int{
				[]int{1, 2, 3, 4},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 2},
				[]int{3, 4},
			}, 4, 1},
			ans{[][]int{
				[]int{1},
				[]int{2},
				[]int{3},
				[]int{4},
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, matrixReshape(p.one, p.r, p.c), "输入:%v", p)
	}
}
