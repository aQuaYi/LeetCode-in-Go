package problem0938

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	L    int
	R    int
	ans  int
}{

	{
		[]int{10, 5, 15, 3, 7, kit.NULL, 18},
		7,
		15,
		32,
	},

	{
		[]int{10, 5, 15, 3, 7, 13, 18, 1, kit.NULL, 6},
		6,
		10,
		23,
	},

	// 可以有多个 testcase
}

func Test_rangeSumBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, rangeSumBST(root, tc.L, tc.R), "输入:%v", tc)
	}
}

func Benchmark_rangeSumBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			rangeSumBST(root, tc.L, tc.R)
		}
	}
}
