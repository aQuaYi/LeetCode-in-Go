package problem0287

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
	one int
}

func Test_Problem0287(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{6, 2, 3, 1, 4, 5, 1},
			},
			ans{
				1,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 6},
			},
			ans{
				6,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findDuplicate(p.nums), "输入:%v", p)
	}
}
