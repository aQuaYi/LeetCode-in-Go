package problem0404

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
		[]int{3, 7, 9, 8, 20, 15, 7},
		[]int{9, 7, 8, 3, 15, 20, 7},
		24,
	},

	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		24,
	},

	// 可以有多个 testcase
}

func Test_sumOfLeftLeaves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, sumOfLeftLeaves(root), "输入:%v", tc)
	}
}

func Benchmark_sumOfLeftLeaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			sumOfLeftLeaves(root)
		}
	}
}
