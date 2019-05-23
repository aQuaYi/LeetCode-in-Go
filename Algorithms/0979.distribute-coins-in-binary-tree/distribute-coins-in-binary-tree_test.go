package problem0979

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
		[]int{3, 0, 0},
		2,
	},

	{
		[]int{0, 3, 0},
		3,
	},

	{
		[]int{1, 0, 2},
		2,
	},

	{
		[]int{1, 0, 0, kit.NULL, 3},
		4,
	},

	// 可以有多个 testcase
}

func Test_distributeCoins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, distributeCoins(root), "输入:%v", tc)
	}
}

func Benchmark_distributeCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			distributeCoins(root)
		}
	}
}
