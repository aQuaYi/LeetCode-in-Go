package problem0108

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0108(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		nums []int
		pre  []int
	}{

		{
			[]int{},
			nil,
		},

		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 1, 3, 6, 5, 7},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.pre, kit.Tree2Preorder(sortedArrayToBST(tc.nums)), "输入:%v", tc)
	}
}
