package problem0814

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

var null = -1 << 63

// tcs is testcase slice
var tcs = []struct {
	root []int
	ans  []int
}{

	{
		[]int{1, null, 0, 0, 1},
		[]int{1, null, 0, null, 1},
	},

	{
		[]int{1, 0, 1, 0, 0, 0, 1},
		[]int{1, null, 1, null, 1},
	},

	{
		[]int{1, 1, 0, 1, 1, 0, 1, 0},
		[]int{1, 1, 0, 1, 1, null, 1},
	},

	// 可以有多个 testcase
}

func Test_pruneTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := kit.Ints2TreeNode(tc.ans)
		get := pruneTree(kit.Ints2TreeNode(tc.root))
		ast.True(ans.Equal(get), "输入:%v", tc)
	}
}

func Benchmark_pruneTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pruneTree(kit.Ints2TreeNode(tc.root))
		}
	}
}
