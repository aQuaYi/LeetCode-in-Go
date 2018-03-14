package problem0216

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
	k int
	n int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0216(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				9,
				45,
			},
			ans{[][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			}},
		},

		question{
			para{
				3,
				7,
			},
			ans{[][]int{
				[]int{1, 2, 4},
			}},
		},

		question{
			para{
				3,
				9,
			},
			ans{[][]int{
				[]int{1, 2, 6},
				[]int{1, 3, 5},
				[]int{2, 3, 4},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, combinationSum3(p.k, p.n), "输入:%v", p)
	}
}
