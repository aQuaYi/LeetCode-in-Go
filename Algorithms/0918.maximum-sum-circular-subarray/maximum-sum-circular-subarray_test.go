package problem0918

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{1, -2, 3, -2},
		3,
	},

	{
		[]int{5, -3, 5},
		10,
	},

	{
		[]int{3, -1, 2, -1},
		4,
	},

	{
		[]int{3, -2, 2, -3},
		3,
	},

	{
		[]int{-2, -3, -1},
		-1,
	},

	// 可以有多个 testcase
}

func Test_maxSubarraySumCircular(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxSubarraySumCircular(tc.A), "输入:%v", tc)
	}
}

func Benchmark_maxSubarraySumCircular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSubarraySumCircular(tc.A)
		}
	}
}
