package problem0951

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root1 []int
	root2 []int
	ans   bool
}{

	{
		[]int{1, 2, 3, 4, 5, 6, kit.NULL, kit.NULL, kit.NULL, 7, 8},
		[]int{1, 3, 2, kit.NULL, 6, 4, 5, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 8, 7},
		true,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, kit.NULL, kit.NULL, kit.NULL, 7, 9},
		[]int{1, 3, 2, kit.NULL, 6, 4, 5, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 8, 7},
		false,
	},

	// 可以有多个 testcase
}

func Test_flipEquiv(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root1 := kit.Ints2TreeNode(tc.root1)
		root2 := kit.Ints2TreeNode(tc.root2)
		ast.Equal(tc.ans, flipEquiv(root1, root2), "输入:%v", tc)
	}
}

func Benchmark_flipEquiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root1 := kit.Ints2TreeNode(tc.root1)
			root2 := kit.Ints2TreeNode(tc.root2)
			flipEquiv(root1, root2)
		}
	}
}
