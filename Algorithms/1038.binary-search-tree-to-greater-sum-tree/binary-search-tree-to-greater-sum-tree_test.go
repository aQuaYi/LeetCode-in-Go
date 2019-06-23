package problem1038

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  []int
}{

	{
		[]int{4, 1, 6, 0, 2, 5, 7, kit.NULL, kit.NULL, kit.NULL, 3, kit.NULL, kit.NULL, kit.NULL, 8},
		[]int{30, 36, 21, 36, 35, 26, 15, kit.NULL, kit.NULL, kit.NULL, 33, kit.NULL, kit.NULL, kit.NULL, 8},
	},

	// 可以有多个 testcase
}

func Test_bstToGst(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ans := kit.Ints2TreeNode(tc.ans)
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(ans, bstToGst(root), "输入:%v", tc)
	}
}

func Benchmark_bstToGst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			bstToGst(root)
		}
	}
}
