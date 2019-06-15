package problem1026

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  int
}{

	{
		[]int{7, 5, 12, 11, 0, 2, kit.NULL, 4, 17, 6, 19, kit.NULL, 16, 18, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 9, 14, 10, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 3, 1, kit.NULL, kit.NULL, 8, kit.NULL, 13, kit.NULL, kit.NULL, 15},
		19,
	},

	{
		[]int{8, 3, 10, 1, 6, kit.NULL, 14, kit.NULL, kit.NULL, 4, 7, 13},
		7,
	},

	// 可以有多个 testcase
}

func Test_maxAncestorDiff(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, maxAncestorDiff(root), "输入:%v", tc)
	}
}

func Benchmark_maxAncestorDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			maxAncestorDiff(root)
		}
	}
}
