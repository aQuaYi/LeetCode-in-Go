package problem0236

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in   []int
	p, q, ans int
}{
	{
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
		5,
		1,
		3,
	},

	{
		[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
		[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
		5,
		4,
		5,
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, que := range questions {
		root := kit.PreIn2Tree(que.pre, que.in)
		p, q := &TreeNode{Val: que.p}, &TreeNode{Val: que.q}
		node := lowestCommonAncestor(root, p, q)
		ast.Equal(que.ans, node.Val)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, que := range questions {
			root := kit.PreIn2Tree(que.pre, que.in)
			p, q := &TreeNode{Val: que.p}, &TreeNode{Val: que.q}
			lowestCommonAncestor(root, p, q)
		}
	}
}
