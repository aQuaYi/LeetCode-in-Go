package Problem0437

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Golang/kit"

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
			[]int{10, 5, 3, 3, -2, 2, 1, -3, 11},
			[]int{3, 3, -2, 5, 2, 1, 10, -3, 11},
			8,
			3,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, pathSum(kit.PreIn2Tree(tc.pre, tc.in), tc.sum), "输入:%v", tc)
	}
}
