package problem0167

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
	numbers []int
	target  int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0167(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{2, 7, 11, 15},
				9,
			},
			ans{
				[]int{1, 2},
			},
		},

		question{
			para{
				[]int{2, 7, 11, 15},
				10,
			},
			ans{
				nil,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, twoSum(p.numbers, p.target), "输入:%v", p)
	}
}
