package problem0236

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

var root = kit.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, kit.NULL, kit.NULL, 7, 4})

var tcs = []struct {
	p, q, ans *TreeNode
}{

	{
		kit.GetTargetNode(root, 5),
		kit.GetTargetNode(root, 1),
		kit.GetTargetNode(root, 3),
	},

	{
		kit.GetTargetNode(root, 4),
		kit.GetTargetNode(root, 7),
		kit.GetTargetNode(root, 2),
	},

	{
		kit.GetTargetNode(root, 4),
		kit.GetTargetNode(root, 5),
		kit.GetTargetNode(root, 5),
	},

	{
		kit.GetTargetNode(root, 5),
		kit.GetTargetNode(root, 4),
		kit.GetTargetNode(root, 5),
	},

	//
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		node := lowestCommonAncestor(root, tc.p, tc.q)
		ast.Equal(tc.ans, node, "p=%d,q=%d", tc.p.Val, tc.q.Val)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lowestCommonAncestor(root, tc.p, tc.q)
		}
	}
}
