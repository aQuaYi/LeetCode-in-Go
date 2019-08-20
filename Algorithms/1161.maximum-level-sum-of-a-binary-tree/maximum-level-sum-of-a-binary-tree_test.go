package problem1161

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
		[]int{1, 7, 0, 7, -8, kit.NULL, kit.NULL},
		2,
	},

	// 可以有多个 testcase
}

func Test_maxLevelSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		root := kit.Ints2TreeNode(tc.root)
		ast.Equal(tc.ans, maxLevelSum(root), "输入:%v", tc)
	}
}

func Benchmark_maxLevelSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.root)
			maxLevelSum(root)
		}
	}
}
