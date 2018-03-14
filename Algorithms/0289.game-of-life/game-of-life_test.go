package problem0289

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
	board [][]int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0289(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{}},
			ans{[][]int{}},
		},

		question{
			para{[][]int{
				[]int{},
			}},
			ans{[][]int{
				[]int{},
			}},
		},

		question{
			para{[][]int{
				[]int{0},
			}},
			ans{[][]int{
				[]int{0},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 1},
			}},
			ans{[][]int{
				[]int{0, 0},
			}},
		},

		question{
			para{[][]int{
				[]int{1},
			}},
			ans{[][]int{
				[]int{0},
			}},
		},

		question{
			para{[][]int{
				[]int{1, 1},
				[]int{1, 0},
			}},
			ans{[][]int{
				[]int{1, 1},
				[]int{1, 1},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		gameOfLife(p.board)
		ast.Equal(a.one, p.board, "输入:%v", p)
	}
}
