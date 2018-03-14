package problem0109

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0109(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		head []int
		ans  []int
	}{

		{
			[]int{},
			nil,
		},

		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, kit.Tree2Inorder(sortedListToBST(kit.Ints2List(tc.head))), "输入:%v", tc)
	}
}
