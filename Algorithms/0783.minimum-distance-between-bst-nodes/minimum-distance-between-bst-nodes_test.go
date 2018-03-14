package problem0783

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{27,34,58,50,44	},
		[]int{27,34,44,50,58},
		6,
	},

	{
		[]int{1, 0, 48, 12, 49},
		[]int{0, 1, 12, 48, 49},
		1,
	},

	{
		[]int{4, 2, 1, 3, 6},
		[]int{1, 2, 3, 4, 6},
		1,
	},

	// 可以有多个 testcase
}

func Test_minDiffInBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, minDiffInBST(root))
	}
}

func Benchmark_minDiffInBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			minDiffInBST(root)
		}
	}
}
