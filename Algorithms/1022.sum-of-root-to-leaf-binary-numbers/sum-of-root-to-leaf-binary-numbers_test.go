package problem1022

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
		[]int{1, kit.NULL, 0},
		2,
	},

	{
		[]int{1, 1},
		3,
	},

	{
		[]int{1, 0, 1, 0, 1, 0, 1},
		22,
	},

	// 可以有多个 testcase
}

func Test_sumRootToLeaf(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, sumRootToLeaf(root), "输入:%v", tc)
	}
}

func Benchmark_sumRootToLeaf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			sumRootToLeaf(root)
		}
	}
}
