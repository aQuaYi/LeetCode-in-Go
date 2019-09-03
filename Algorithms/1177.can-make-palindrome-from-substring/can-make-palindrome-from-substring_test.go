package problem1177

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s       string
	queries [][]int
	ans     []bool
}{

	{
		"abcda",
		[][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}},
		[]bool{true, false, false, true, true},
	},

	// 可以有多个 testcase
}

func Test_canMakePaliQueries(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, canMakePaliQueries(tc.s, tc.queries), "输入:%v", tc)
	}
}

func Benchmark_canMakePaliQueries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canMakePaliQueries(tc.s, tc.queries)
		}
	}
}
