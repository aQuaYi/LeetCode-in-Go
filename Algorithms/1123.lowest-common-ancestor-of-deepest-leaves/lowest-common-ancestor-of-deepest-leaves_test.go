package problem1123

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
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	},

	{
		[]int{1, 2, 3, kit.NULL, kit.NULL, 4},
		[]int{4},
	},

	{
		[]int{1, 2, 3, 4},
		[]int{4},
	},

	{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_lcaDeepestLeaves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ans := kit.Tree2ints(lcaDeepestLeaves(root))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_lcaDeepestLeaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			lcaDeepestLeaves(root)
		}
	}
}
