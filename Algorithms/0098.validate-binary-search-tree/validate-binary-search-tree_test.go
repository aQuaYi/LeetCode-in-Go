package problem0098

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0098(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     bool
	}{

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			false,
		},

		{
			[]int{},
			[]int{},
			true,
		},

		{
			[]int{2, 1, 3},
			[]int{1, 2, 3},
			true,
		},

		{
			[]int{10, 5, 15, 6, 20},
			[]int{5, 10, 6, 15, 20},
			false,
		},

		{
			[]int{3, 2, 1},
			[]int{2, 3, 1},
			false,
		},
		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isValidBST(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
