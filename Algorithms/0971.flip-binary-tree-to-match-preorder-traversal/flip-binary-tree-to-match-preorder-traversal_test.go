package problem0971

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root   []int
	voyage []int
	ans    []int
}{

	{
		[]int{1, 2},
		[]int{1, 2},
		[]int{},
	},

	{
		[]int{1, kit.NULL, 2},
		[]int{1, 2},
		[]int{},
	},

	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{1},
	},

	{
		[]int{1, 2},
		[]int{2, 1},
		[]int{-1},
	},

	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{},
	},

	// 可以有多个 testcase
}

func Test_flipMatchVoyage(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, flipMatchVoyage(root, tc.voyage), "输入:%v", tc)
	}
}

func Benchmark_flipMatchVoyage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			flipMatchVoyage(root, tc.voyage)
		}
	}
}
