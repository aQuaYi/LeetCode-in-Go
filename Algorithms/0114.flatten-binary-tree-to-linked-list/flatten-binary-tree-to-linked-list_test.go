package problem0114

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0114(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in, post []int
	}{

		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]int{3, 2, 1},
		},

		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 2, 4, 1, 5, 6},
			[]int{6, 5, 4, 3, 2, 1},
		},

		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.PreIn2Tree(tc.pre, tc.in)
		flatten(root)
		ast.Equal(tc.post, kit.Tree2Postorder(root), "输入:%v", tc)
	}
}
