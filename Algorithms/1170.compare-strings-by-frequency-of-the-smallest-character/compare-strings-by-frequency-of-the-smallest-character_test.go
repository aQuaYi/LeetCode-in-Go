package problem1170

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	queries []string
	words   []string
	ans     []int
}{

	{
		[]string{"cbd"},
		[]string{"zaaaz"},
		[]int{1},
	},

	{
		[]string{"bbb", "cc"},
		[]string{"a", "aa", "aaa", "aaaa"},
		[]int{1, 2},
	},

	// 可以有多个 testcase
}

func Test_numSmallerByFrequency(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, numSmallerByFrequency(tc.queries, tc.words), "输入:%v", tc)
	}
}

func Benchmark_numSmallerByFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSmallerByFrequency(tc.queries, tc.words)
		}
	}
}
