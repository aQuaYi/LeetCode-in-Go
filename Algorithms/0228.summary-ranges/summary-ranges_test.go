package problem0228

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
	nums []int
}

// ans 是答案
type ans struct {
	one []string
}

func Test_Problem0228(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{},
			},
			ans{
				[]string{},
			},
		},

		question{
			para{
				[]int{0, 1, 2, 4, 5, 7},
			},
			ans{
				[]string{"0->2", "4->5", "7"},
			},
		},

		question{
			para{
				[]int{0, 2, 3, 4, 6, 8, 9},
			},
			ans{
				[]string{"0", "2->4", "6", "8->9"},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, summaryRanges(p.nums), "输入:%v", p)
	}
}
