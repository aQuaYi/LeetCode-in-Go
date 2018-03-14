package problem0100

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	aPre []int
	aIn  []int
	bPre []int
	bIn  []int
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0100(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3},
				[]int{2, 1, 3},
				[]int{1, 2, 3},
				[]int{2, 1, 3},
			},
			ans{
				true,
			},
		},

		question{
			para{
				[]int{1, 2, 3},
				[]int{2, 1, 3},
				[]int{1, 2, 3},
				[]int{1, 2, 3},
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

		ast.Equal(a.one, isSameTree(kit.PreIn2Tree(p.aPre, p.aIn), kit.PreIn2Tree(p.bPre, p.bIn)), "输入:%v", p)
	}
}
