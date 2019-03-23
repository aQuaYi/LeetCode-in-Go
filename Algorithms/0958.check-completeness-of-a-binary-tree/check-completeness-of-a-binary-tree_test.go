package problem0958

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
		[]int{1, 2, 3, 4, 5, 6},
		true,
	},

	{
		[]int{1, 2, 3, 4, 5, kit.NULL, 7},
		false,
	},

	// 可以有多个 testcase
}

func Test_isCompleteTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, isCompleteTree(root), "输入:%v", tc)
	}
}

func Benchmark_isCompleteTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			isCompleteTree(root)
		}
	}
}
