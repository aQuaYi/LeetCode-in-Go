package problem1207

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{1, 2, 2, 1, 1, 3},
		true,
	},

	{
		[]int{1, 2},
		false,
	},

	{
		[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		true,
	},

	// 可以有多个 testcase
}

func Test_uniqueOccurrences(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, uniqueOccurrences(tc.A), "输入:%v", tc)
	}
}

func Benchmark_uniqueOccurrences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniqueOccurrences(tc.A)
		}
	}
}
