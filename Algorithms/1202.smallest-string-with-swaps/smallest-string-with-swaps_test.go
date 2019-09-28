package problem1202

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s     string
	pairs [][]int
	ans   string
}{

	{
		"dcab",
		[][]int{{0, 3}, {1, 2}},
		"bacd",
	},

	{
		"dcab",
		[][]int{{0, 3}, {1, 2}, {0, 2}},
		"abcd",
	},

	{
		"cba",
		[][]int{{0, 1}, {1, 2}},
		"abc",
	},

	// 可以有多个 testcase
}

func Test_smallestStringWithSwaps(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, smallestStringWithSwaps(tc.s, tc.pairs), "输入:%v", tc)
	}
}

func Benchmark_smallestStringWithSwaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestStringWithSwaps(tc.s, tc.pairs)
		}
	}
}
