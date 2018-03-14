package problem0667

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
	one []int
}

func Test_Problem0667(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				3,
				1,
			},
			ans{
				[]int{1, 2, 3},
			},
		},

		question{
			para{
				3,
				2,
			},
			ans{
				[]int{3, 1, 2},
			},
		},

		question{
			para{
				5,
				2,
			},
			ans{
				[]int{5, 1, 2, 3, 4},
			},
		},

		question{
			para{
				5,
				3,
			},
			ans{
				[]int{1, 5, 2, 3, 4},
			},
		},

		question{
			para{
				5,
				4,
			},
			ans{
				[]int{5, 1, 4, 2, 3},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, constructArray(p.n, p.k), "输入:%v", p)
	}
}
