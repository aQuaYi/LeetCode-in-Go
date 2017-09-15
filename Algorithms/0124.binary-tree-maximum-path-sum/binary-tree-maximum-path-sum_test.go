package Problem0124

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Golang/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0124(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     int
	}{

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			6,
		},

		{
			[]int{4, 2, 1, 3, 6, 5, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			22,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, maxPathSum(root), "输入:%v", tc)
	}
}
