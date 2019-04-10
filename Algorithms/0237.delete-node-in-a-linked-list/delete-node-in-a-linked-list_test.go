package problem0237

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	node int
	ans  []int
}{

	{
		[]int{4, 5, 1, 9},
		5,
		[]int{4, 1, 9},
	},

	{
		[]int{4, 5, 1, 9},
		1,
		[]int{4, 5, 9},
	},

	// 可以有多个 testcase
}

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := kit.Ints2List(tc.ints)
		node := head.GetNodeWith(tc.node)
		deleteNode(node)
		ast.Equal(tc.ans, kit.List2Ints(head), "输入:%v", tc)
	}
}

func Benchmark_deleteNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.ints)
			node := head.GetNodeWith(tc.node)
			deleteNode(node)
		}
	}
}
