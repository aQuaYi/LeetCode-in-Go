package problem0073

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
type para struct {
	matrix [][]int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0073(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{
				[]int{1, 0, 3, 4},
				[]int{5, 0, 7, 0},
				[]int{9, 1, 1, 1},
			}},
			ans{[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{9, 0, 1, 0},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 2, 3, 4},
				[]int{0, 2, 3, 4},
				[]int{9, 1, 1, 1},
			}},
			ans{[][]int{
				[]int{0, 2, 3, 4},
				[]int{0, 0, 0, 0},
				[]int{0, 1, 1, 1},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 0, 7, 0},
				[]int{9, 1, 1, 1},
			}},
			ans{[][]int{
				[]int{1, 0, 3, 0},
				[]int{0, 0, 0, 0},
				[]int{9, 0, 1, 0},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		setZeroes(p.matrix)
		ast.Equal(a.one, p.matrix, "输入:%v", p)
	}
}
