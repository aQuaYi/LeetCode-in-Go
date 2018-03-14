package problem0543

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{},
		[]int{},
		0,
	},

	{
		[]int{1, 2, 4, 5, 3},
		[]int{4, 2, 5, 1, 3},
		3,
	},

	// 可以有多个 testcase
}

func Test_diameterOfBinaryTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, diameterOfBinaryTree(root), "输入:%v", tc)
	}
}

func Benchmark_diameterOfBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			diameterOfBinaryTree(root)
		}
	}
}
