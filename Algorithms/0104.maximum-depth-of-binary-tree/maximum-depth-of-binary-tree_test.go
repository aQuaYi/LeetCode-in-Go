package problem0104

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0104(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     int
	}{

		{
			[]int{},
			[]int{},
			0,
		},

		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			3,
		},

		{
			[]int{3, 9, 20, 15, 10, 7},
			[]int{9, 3, 10, 15, 20, 7},
			4,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxDepth(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
