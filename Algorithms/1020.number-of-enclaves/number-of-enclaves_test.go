package problem1020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   [][]int
	ans int
}{

	{
		[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
		3,
	},

	{
		[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
		0,
	},

	// 可以有多个 testcase
}

func Test_numEnclaves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numEnclaves(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numEnclaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numEnclaves(tc.A)
		}
	}
}
