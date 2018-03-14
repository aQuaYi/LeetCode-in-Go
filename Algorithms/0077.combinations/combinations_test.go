package problem0077

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
	n int
	k int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0077(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				4,
				2,
			},
			ans{
				[][]int{
					[]int{1, 2},
					[]int{1, 3},
					[]int{1, 4},
					[]int{2, 3},
					[]int{2, 4},
					[]int{3, 4},
				},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, combine(p.n, p.k), "输入:%v", p)
	}
}
