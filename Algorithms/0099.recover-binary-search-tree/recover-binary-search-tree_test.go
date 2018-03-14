package problem0099

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0099(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     []int
	}{

		{
			[]int{3, 2, 1},
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},

		{
			[]int{6, 2, 1, 3, 4, 5, 7},
			[]int{1, 2, 3, 6, 5, 4, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},

		{
			[]int{5, 2, 1, 3, 6, 4, 7},
			[]int{1, 2, 3, 5, 4, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},

		{
			[]int{3, 2, 1, 4, 6, 5, 7},
			[]int{1, 2, 4, 3, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},

		{
			[]int{4, 2, 1, 5, 6, 3, 7},
			[]int{1, 2, 5, 4, 3, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.PreIn2Tree(tc.pre, tc.in)
		recoverTree(root)
		ast.Equal(tc.ans, kit.Tree2Inorder(root), "输入:%v", tc)
	}
}
