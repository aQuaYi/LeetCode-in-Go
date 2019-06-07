package problem1008

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	preorder []int
	ans      []int
}{

	{
		[]int{8, 5, 1, 7, 10, 12},
		[]int{8, 5, 10, 1, 7, kit.NULL, 12},
	},

	// 可以有多个 testcase
}

func Test_bstFromPreorder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ans := kit.Tree2ints(bstFromPreorder(tc.preorder))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_bstFromPreorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			bstFromPreorder(tc.preorder)
		}
	}
}
