package problem0949

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans string
}{

	{
		[]int{1, 2, 3, 4},
		"23:41",
	},

	{
		[]int{5, 5, 5, 5},
		"",
	},

	// 可以有多个 testcase
}

func Test_largestTimeFromDigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, largestTimeFromDigits(tc.A), "输入:%v", tc)
	}
}

func Benchmark_largestTimeFromDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestTimeFromDigits(tc.A)
		}
	}
}
