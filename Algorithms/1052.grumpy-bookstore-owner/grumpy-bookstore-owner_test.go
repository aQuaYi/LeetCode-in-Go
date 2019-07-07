package problem1052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	customers []int
	grumpy    []int
	X         int
	ans       int
}{

	{
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1},
		3,
		16,
	},

	// 可以有多个 testcase
}

func Test_maxSatisfied(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSatisfied(tc.customers, tc.grumpy, tc.X), "输入:%v", tc)
	}
}

func Benchmark_maxSatisfied(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSatisfied(tc.customers, tc.grumpy, tc.X)
		}
	}
}
