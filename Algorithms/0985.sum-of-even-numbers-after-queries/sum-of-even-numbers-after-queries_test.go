package problem0985

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A       []int
	queries [][]int
	ans     []int
}{

	{
		[]int{1, 2, 3, 4},
		[][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
		[]int{8, 6, 2, 4},
	},

	// 可以有多个 testcase
}

func Test_sumEvenAfterQueries(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, sumEvenAfterQueries(tc.A, tc.queries), "输入:%v", tc)
	}
}

func Benchmark_sumEvenAfterQueries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumEvenAfterQueries(tc.A, tc.queries)
		}
	}
}
