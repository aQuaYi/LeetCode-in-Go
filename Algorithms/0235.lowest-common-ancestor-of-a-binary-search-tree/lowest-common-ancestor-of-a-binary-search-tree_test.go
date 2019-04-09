package problem0235

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
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
		8,
		6,
	},

	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		2,
		4,
		2,
	},

	{
		[]int{6, 2, 0, 4, 3, 5, 8, 7, 9},
		[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
		7,
		9,
		8,
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)
	for _, que := range questions {
		root := kit.PreIn2Tree(que.pre, que.in)

		p := &TreeNode{
			Val: que.p,
		}

		q := &TreeNode{
			Val: que.q,
		}

		node := lowestCommonAncestor(root, p, q)
		ast.Equal(que.ans, node.Val)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, que := range questions {
			root := kit.PreIn2Tree(que.pre, que.in)
			p := &TreeNode{
				Val: que.p,
			}
			q := &TreeNode{
				Val: que.q,
			}
			lowestCommonAncestor(root, p, q)
		}
	}
}
