package problem0968

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
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		5,
	},

	{
		[]int{0, 0, 0, 0, 0, kit.NULL, kit.NULL, 0},
		3,
	},

	{
		[]int{0, 0, 0, 0, 0, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 0, 0},
		3,
	},

	{
		[]int{0},
		1,
	},

	{
		[]int{0, 0, kit.NULL, 0, 0},
		1,
	},

	{
		[]int{0, 0, kit.NULL, 0, kit.NULL, 0, kit.NULL, kit.NULL, 0},
		2,
	},

	// 可以有多个 testcase
}

func Test_minCameraCover(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, minCameraCover(root), "输入:%v", tc)
	}
}

func Benchmark_minCameraCover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			minCameraCover(root)
		}
	}
}
