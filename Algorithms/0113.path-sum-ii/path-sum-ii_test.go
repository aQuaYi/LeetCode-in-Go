package problem0113

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0113(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		sum     int
		ans     [][]int
	}{

		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 5, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 5, 4, 1},
			22,
			[][]int{
				[]int{5, 4, 11, 2},
				[]int{5, 8, 4, 5},
			},
		},

		{
			[]int{5},
			[]int{5},
			5,
			[][]int{
				[]int{5},
			},
		},

		{
			[]int{5},
			[]int{5},
			4,
			[][]int{},
		},

		{
			[]int{},
			[]int{},
			4,
			[][]int{},
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, pathSum(kit.PreIn2Tree(tc.pre, tc.in), tc.sum), "输入:%v", tc)
	}
}
