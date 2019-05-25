package problem0982

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
		[]int{2, 1, 3},
		12,
	},

	// 可以有多个 testcase
}

func Test_countTriplets(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, countTriplets(tc.A), "输入:%v", tc)
	}
}

func Benchmark_countTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countTriplets(tc.A)
		}
	}
}
