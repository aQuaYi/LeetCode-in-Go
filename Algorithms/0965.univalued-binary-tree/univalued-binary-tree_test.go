package problem0965

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  bool
}{

	{
		[]int{1, 1, 1, 1, 1, kit.NULL, 1},
		true,
	},

	{
		[]int{2, 2, 2, 5, 2},
		false,
	},

	// 可以有多个 testcase
}

func Test_isUnivalTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, isUnivalTree(root), "输入:%v", tc)
	}
}

func Benchmark_isUnivalTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			isUnivalTree(root)
		}
	}
}
