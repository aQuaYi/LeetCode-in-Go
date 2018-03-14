package problem0530

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
		[]int{1, 3, 2},
		[]int{1, 2, 3},
		1,
	},

	// 可以有多个 testcase
}

func Test_getMinimumDifference(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, getMinimumDifference(root), "输入:%v", tc)
	}
}

func Benchmark_getMinimumDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			getMinimumDifference(root)
		}
	}
}
