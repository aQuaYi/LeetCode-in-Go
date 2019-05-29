package problem0986

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   [][]int
	B   [][]int
	ans [][]int
}{

	{
		[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 2}, {8, 12}, {15, 24}, {25, 26}},
		[][]int{{1, 2}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
	},

	{
		[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
	},

	{
		[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 1}, {8, 12}, {15, 24}, {25, 26}},
		[][]int{{1, 1}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
	},

	// 可以有多个 testcase
}

func Test_intervalIntersection(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, intervalIntersection(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_intervalIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intervalIntersection(tc.A, tc.B)
		}
	}
}
