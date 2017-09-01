package Problem0188

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
	k      int
	prices []int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0188(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				3,
				[]int{3, 7, 2, 5, 7, 5, 5, 6, 7, 7, 2, 2, 5, 7, 12, 17, 7, 1, 5, 7},
			},
			ans{
				26,
			},
		},

		question{
			para{
				2,
				[]int{},
			},
			ans{
				0,
			},
		},

		question{
			para{
				3,
				[]int{3, 7, 2, 5, 7, 5, 2, 25, 28, 7, 7, 2, 4, 5, 67, 7, 7, 2, 2, 5, 7, 12, 17, 712, 2, 62, 6, 2, 5, 6, 2, 5, 1, 7, 22, 1, 5, 7},
			},
			ans{
				835,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxProfit(p.k, p.prices), "输入:%v", p)
	}
}
