package problem1128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dominoes [][]int
	ans      int
}{

	{
		[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
		1,
	},

	// 可以有多个 testcase
}

func Test_numEquivDominoPairs(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numEquivDominoPairs(tc.dominoes), "输入:%v", tc)
	}
}

func Benchmark_numEquivDominoPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numEquivDominoPairs(tc.dominoes)
		}
	}
}
