package problem0095

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
	n int
}

// ans 是答案
type ans struct {
	in []int
}

func Test_Problem0095(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				3,
			},
			ans{
				[]int{1, 2, 3},
			},
		},

		question{
			para{
				4,
			},
			ans{
				[]int{1, 2, 3, 4},
			},
		},

		question{
			para{
				5,
			},
			ans{
				[]int{1, 2, 3, 4, 5},
			},
		},

		question{
			para{
				6,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		roots := generateTrees(p.n)
		for _, root := range roots {
			ast.Equal(a.in, kit.Tree2Inorder(root), "输入:%d", p.n)
		}
	}

	ast.Nil(generateTrees(0), "generateTrees(0) == nil")
}

func Test_indexOf(t *testing.T) {
	ast := assert.New(t)

	ast.Panics(func() { indexOf(3, []int{0, 1, 2}) })
}
