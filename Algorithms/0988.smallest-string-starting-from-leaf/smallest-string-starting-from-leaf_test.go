package problem0988

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  string
}{

	{

		[]int{25, 1, kit.NULL, 0, 0, 1, kit.NULL, kit.NULL, kit.NULL, 0},
		"ababz",
	},

	{
		[]int{4, 0, 1, 1},
		"bae",
	},

	{
		[]int{25, 1, 3, 1, 3, 0, 2},
		"adz",
	},

	{
		[]int{0, 1, 2, 3, 4, 3, 4},
		"dba",
	},

	{
		[]int{25, 1, 3, 1, 3, 0, 2},
		"adz",
	},

	{
		[]int{2, 2, 1, kit.NULL, 1, 0, kit.NULL, 0},
		"abc",
	},

	// 可以有多个 testcase
}

func Test_smallestFromLeaf(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, smallestFromLeaf(root), "输入:%v", tc)
	}
}

func Benchmark_smallestFromLeaf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			smallestFromLeaf(root)
		}
	}
}
