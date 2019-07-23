package problem1104

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	label int
	ans   []int
}{

	{
		14,
		[]int{1, 3, 4, 14},
	},

	{
		26,
		[]int{1, 2, 6, 10, 26},
	},

	// 可以有多个 testcase
}

func Test_pathInZigZagTree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, pathInZigZagTree(tc.label), "输入:%v", tc)
	}
}

func Benchmark_pathInZigZagTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pathInZigZagTree(tc.label)
		}
	}
}
