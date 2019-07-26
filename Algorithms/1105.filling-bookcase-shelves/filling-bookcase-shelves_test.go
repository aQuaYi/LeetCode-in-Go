package problem1105

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	books      [][]int
	shelfWidth int
	ans        int
}{

	{
		[][]int{{1, 1}, {3, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
		4,
		8,
	},

	{
		[][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
		4,
		6,
	},

	// 可以有多个 testcase
}

func Test_minHeightShelves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minHeightShelves(tc.books, tc.shelfWidth), "输入:%v", tc)
	}
}

func Benchmark_minHeightShelves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minHeightShelves(tc.books, tc.shelfWidth)
		}
	}
}
