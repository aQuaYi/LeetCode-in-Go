package problem0814

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

var null = -1 << 63

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  *TreeNode
}{

	{
		kit.Ints2TreeNode([]int{1, null, 0, 0, 1}),
		kit.Ints2TreeNode([]int{1, null, 0, null, 1}),
	},

	{
		kit.Ints2TreeNode([]int{1, 0, 1, 0, 0, 0, 1}),
		kit.Ints2TreeNode([]int{1, null, 1, null, 1}),
	},

	{
		kit.Ints2TreeNode([]int{1, 1, 0, 1, 1, 0, 1, 0}),
		kit.Ints2TreeNode([]int{1, 1, 0, 1, 1, null, 1}),
	},

	// 可以有多个 testcase
}

func Test_pruneTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.True(tc.ans.Equal(pruneTree(tc.root)), "输入:%v", tc)
	}
}

func Benchmark_pruneTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pruneTree(tc.root)
		}
	}
}
