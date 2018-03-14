package problem0437

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0437(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		sum     int
		ans     int
	}{

		{
			[]int{8, 5, 3, 5, -2, 2, 1, -3, 11, 0, 8},
			[]int{5, 3, -2, 5, 2, 1, 8, -3, 11, 0, 8},
			8,
			8,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, pathSum(kit.PreIn2Tree(tc.pre, tc.in), tc.sum), "输入:%v", tc)
	}
}
