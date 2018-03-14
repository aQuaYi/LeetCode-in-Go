package problem0039

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
	candidates []int
	target     int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0039(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{7, 3, 2}, 18},
			ans{[][]int{
				[]int{2, 2, 2, 2, 2, 2, 2, 2, 2},
				[]int{2, 2, 2, 2, 2, 2, 3, 3},
				[]int{2, 2, 2, 2, 3, 7},
				[]int{2, 2, 2, 3, 3, 3, 3},
				[]int{2, 2, 7, 7},
				[]int{2, 3, 3, 3, 7},
				[]int{3, 3, 3, 3, 3, 3},
			}},
		},

		question{
			para{[]int{2, 3, 6, 7}, 7},
			ans{[][]int{
				[]int{2, 2, 3},
				[]int{7},
			}},
		},

		question{
			para{[]int{8, 7, 4, 3}, 11},
			ans{[][]int{
				[]int{3, 4, 4},
				[]int{3, 8},
				[]int{4, 7},
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, combinationSum(p.candidates, p.target), "输入:%v", p)
	}
}
