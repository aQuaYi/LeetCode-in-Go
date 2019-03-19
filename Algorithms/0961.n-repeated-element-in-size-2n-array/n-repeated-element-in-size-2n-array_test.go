package problem0961

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{1, 2, 3, 3},
		3,
	},

	{
		[]int{2, 1, 2, 5, 3, 2},
		2,
	},

	{
		[]int{5, 1, 5, 2, 5, 3, 5, 4},
		5,
	},

	// 可以有多个 testcase
}

func Test_repeatedNTimes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, repeatedNTimes(tc.A), "输入:%v", tc)
	}
}

func Benchmark_repeatedNTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			repeatedNTimes(tc.A)
		}
	}
}
