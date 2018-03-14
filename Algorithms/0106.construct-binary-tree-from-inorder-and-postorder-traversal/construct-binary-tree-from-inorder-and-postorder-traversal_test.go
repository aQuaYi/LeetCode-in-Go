package problem0106

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
	inorder   []int
	postorder []int
}

// ans 是答案
type ans struct {
	pre []int
}

func Test_Problem0106(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{4, 2, 5, 1, 6, 3, 7},
				[]int{4, 5, 2, 6, 7, 3, 1},
			},
			ans{
				[]int{1, 2, 4, 5, 3, 6, 7},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.pre, kit.Tree2Preorder(buildTree(p.inorder, p.postorder)), "输入:%v", p)
	}

	ast.Nil(buildTree([]int{}, []int{}))

	ast.Equal(0, indexOf(3, []int{0, 1, 2}))
}
