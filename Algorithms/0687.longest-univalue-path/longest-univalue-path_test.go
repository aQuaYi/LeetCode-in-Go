package problem0687

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

const null = -1 << 63

// tcs is testcase slice
var tcs = []struct {
	ints []int
	ans  int
}{

	{
		[]int{},
		0,
	},

	{
		[]int{1, 4, 5, 4, 4, null, 5},
		2,
	},

	{
		[]int{5, 4, 5, 1, 1, null, 5},
		2,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2TreeNode(tc.ints)
		ast.Equal(tc.ans, longestUnivaluePath(root), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.ints)
			longestUnivaluePath(root)
		}
	}
}
