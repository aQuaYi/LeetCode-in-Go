package problem0110

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0110(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     bool
	}{

		{
			[]int{2, 1, 5, 4, 3, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			false,
		},

		{
			[]int{4, 2, 1, 3, 6, 5, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isBalanced(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
