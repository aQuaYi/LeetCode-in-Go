package problem0209

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
	s    int
	nums []int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0209(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				15,
				[]int{1, 2, 3, 4, 5},
			},
			ans{
				5,
			},
		},

		question{
			para{
				700,
				[]int{2, 3, 1, 2, 4, 3},
			},
			ans{
				0,
			},
		},

		question{
			para{
				7,
				[]int{2, 3, 1, 2, 4, 3},
			},
			ans{
				2,
			},
		},

		question{
			para{
				7,
				[]int{},
			},
			ans{
				0,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, minSubArrayLen(p.s, p.nums), "输入:%v", p)
	}
}
