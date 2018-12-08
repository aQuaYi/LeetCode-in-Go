package problem0897

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

var null = kit.NULL

// tcs is testcase slice
var tcs = []struct {
	p   []int
	ans []int
}{

	{
		[]int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9},
		[]int{1, null, 2, null, 3, null, 4, null, 5, null, 6, null, 7, null, 8, null, 9},
	},

	// 可以有多个 testcase
}

func Test_increasingBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.Ints2TreeNode(tc.p)
		ans := kit.Tree2ints(increasingBST(root))
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_myFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.Ints2TreeNode(tc.p)
			increasingBST(root)
		}
	}
}
