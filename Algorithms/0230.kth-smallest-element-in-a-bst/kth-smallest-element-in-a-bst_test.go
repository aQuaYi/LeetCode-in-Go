package problem0230

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	k       int
	ans     int
}{

	{
		[]int{1, 2},
		[]int{1, 2},
		2,
		2,
	},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		1,
		1,
	},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		2,
		2,
	},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		3,
		3,
	},

	// 可以有多个 testcase
}

func Test_kthSmallest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, kthSmallest(root, tc.k), "输入:%v", tc)
	}
}

func Benchmark_kthSmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			kthSmallest(root, tc.k)
		}
	}
}
