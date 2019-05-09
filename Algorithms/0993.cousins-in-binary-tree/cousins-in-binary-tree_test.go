package problem0993

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	x    int
	y    int
	ans  bool
}{

	{
		[]int{1, 2, 3, 4},
		4,
		3,
		false,
	},

	{
		[]int{1, 2, 3, kit.NULL, 4, kit.NULL, 5},
		5,
		4,
		true,
	},

	{
		[]int{1, 2, 3, kit.NULL, 4},
		2,
		3,
		false,
	},

	// 可以有多个 testcase
}

func Test_isCousins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, isCousins(root, tc.x, tc.y), "输入:%v", tc)
	}
}

func Benchmark_isCousins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			isCousins(root, tc.x, tc.y)
		}
	}
}
