package problem0662

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
		[]int{1, 3, 5, 4, 2, 9},
		[]int{5, 3, 4, 1, 2, 9},
		4,
	},

	{
		[]int{1, 3, 5, 4},
		[]int{5, 3, 4, 1},
		2,
	},

	{
		[]int{1, 3, 5, 2},
		[]int{5, 3, 1, 2},
		2,
	},

	{
		[]int{1, 3, 5, 6, 2, 9, 7},
		[]int{6, 5, 3, 1, 2, 9, 7},
		8,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, widthOfBinaryTree(root), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			widthOfBinaryTree(root)
		}
	}
}
