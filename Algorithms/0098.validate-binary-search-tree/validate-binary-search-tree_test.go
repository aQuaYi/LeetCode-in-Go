package Problem0098

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Golang/kit"

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
			[]int{2, 1, 3},
			[]int{1, 2, 3},
			true,
		},

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			false,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isValidBST(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
