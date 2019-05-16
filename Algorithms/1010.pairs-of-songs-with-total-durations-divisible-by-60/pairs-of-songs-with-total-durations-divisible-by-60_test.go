package problem1010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	time []int
	ans  int
}{

	{
		[]int{30, 20, 150, 100, 40},
		3,
	},

	{
		[]int{60, 60, 60},
		3,
	},

	// 可以有多个 testcase
}

func Test_numPairsDivisibleBy60(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numPairsDivisibleBy60(tc.time), "输入:%v", tc)
	}
}

func Benchmark_numPairsDivisibleBy60(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numPairsDivisibleBy60(tc.time)
		}
	}
}
