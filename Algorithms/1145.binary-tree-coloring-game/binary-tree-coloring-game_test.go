package problem1145

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root []int
	n    int
	x    int
	ans  bool
}{

	{
		[]int{6, 3, kit.NULL, 7, 4, kit.NULL, kit.NULL, kit.NULL, 2, kit.NULL, 1, kit.NULL, 5},
		7,
		3,
		true,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		11,
		3,
		true,
	},

	{
		[]int{3, kit.NULL, 6, kit.NULL, 7, 4, kit.NULL, 1, 2, kit.NULL, kit.NULL, 5},
		7,
		4,
		false,
	},

	// 可以有多个 testcase
}

func Test_btreeGameWinningMove(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, btreeGameWinningMove(root, tc.n, tc.x), "输入:%v", tc)
	}
}

func Benchmark_btreeGameWinningMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			btreeGameWinningMove(root, tc.n, tc.x)
		}
	}
}
