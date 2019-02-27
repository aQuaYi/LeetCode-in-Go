package problem0942

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []int
}{

	{
		"IDID",
		[]int{0, 4, 1, 3, 2},
	},

	{
		"III",
		[]int{0, 1, 2, 3},
	},

	{
		"DDI",
		[]int{3, 2, 0, 1},
	},

	// 可以有多个 testcase
}

func Test_diStringMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, diStringMatch(tc.S), "输入:%v", tc)
	}
}

func Benchmark_diStringMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			diStringMatch(tc.S)
		}
	}
}
