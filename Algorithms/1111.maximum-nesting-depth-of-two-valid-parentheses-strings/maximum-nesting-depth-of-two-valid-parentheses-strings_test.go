package problem1111

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	seq string
	ans []int
}{

	{
		"(()())",
		[]int{0, 1, 1, 1, 1, 0},
	},

	{
		"()(())()",
		[]int{0, 0, 0, 1, 1, 0, 0, 0},
	},

	// 可以有多个 testcase
}

func Test_maxDepthAfterSplit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxDepthAfterSplit(tc.seq), "输入:%v", tc)
	}
}

func Benchmark_maxDepthAfterSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxDepthAfterSplit(tc.seq)
		}
	}
}
