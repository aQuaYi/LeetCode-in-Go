package problem0219

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
	k    int
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0219(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 1, 3, 4, 5},
				2,
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				0,
			},
			ans{
				false,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{
				false,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, containsNearbyDuplicate(p.nums, p.k), "输入:%v", p)
	}
}
