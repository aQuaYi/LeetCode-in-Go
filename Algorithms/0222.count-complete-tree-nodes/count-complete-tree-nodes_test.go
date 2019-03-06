package problem0222

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	count   int
}{
	{
		[]int{1, 2, 4, 5, 3, 6},
		[]int{4, 2, 5, 1, 6, 3},
		6,
	},
}

func Test_countNodes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := prein2Tree(q.pre, q.in)
		ast.Equal(q.count, countNodes(root))
	}
}

func Benchmark_countNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := prein2Tree(q.pre, q.in)
			countNodes(root)
		}
	}
}

func prein2Tree(pre, in []int) *TreeNode {
	preLen := len(pre)
	inLen := len(in)
	if preLen != inLen {
		panic("pre in 数组长度不一致!")
	}

	if preLen == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if preLen == 1 {
		return root
	}

	idx := indexOf(root.Val, in)
	if idx == -1 {
		panic("pre 或 in 数组有错!")
	}

	root.Left = prein2Tree(pre[1:idx+1], in[:idx])
	root.Right = prein2Tree(pre[idx+1:], in[idx+1:])

	return root
}

func indexOf(val int, nums []int) int {
	for k, v := range nums {
		if v == val {
			return k
		}
	}

	return -1
}
