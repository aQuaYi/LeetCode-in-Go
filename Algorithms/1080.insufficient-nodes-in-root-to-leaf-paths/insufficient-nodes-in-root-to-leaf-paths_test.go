package problem1080

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root  []int
	limit int
	ans   []int
}{

	{
		[]int{1, 2, -3, -5, kit.NULL, 4, kit.NULL},
		-1,
		[]int{1, kit.NULL, -3, 4},
	},

	{
		[]int{1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14},
		1,
		[]int{1, 2, 3, 4, kit.NULL, kit.NULL, 7, 8, 9, kit.NULL, 14},
	},

	{
		[]int{5, 4, 8, 11, kit.NULL, 17, 4, 7, 1, kit.NULL, kit.NULL, 5, 3},
		22,
		[]int{5, 4, 8, 11, kit.NULL, 17, 4, 7, kit.NULL, kit.NULL, kit.NULL, 5},
	},

	{
		[]int{1, 2, -3, -5, kit.NULL, 4, kit.NULL},
		100,
		[]int{},
	},

	// 可以有多个 testcase
}

func Test_sufficientSubset(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, kit.Tree2ints(sufficientSubset(root, tc.limit)), "输入:%v", tc)
	}
}

func Benchmark_sufficientSubset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			sufficientSubset(root, tc.limit)
		}
	}
}
