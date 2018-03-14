package problem0040

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
			para{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			ans{[][]int{
				[]int{1, 1, 6},
				[]int{1, 2, 5},
				[]int{1, 7},
				[]int{2, 6},
			}},
		},

		question{
			para{[]int{1}, 1},
			ans{[][]int{
				[]int{1},
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, combinationSum2(p.candidates, p.target), "输入:%v", p)
	}
}
