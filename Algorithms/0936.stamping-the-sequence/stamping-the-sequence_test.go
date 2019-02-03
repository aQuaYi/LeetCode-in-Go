package problem0936

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stamp  string
	target string
	ans    []int
}{

	{
		"abc",
		"ababc",
		[]int{0, 2},
	},

	{
		"abca",
		"aabcaca",
		[]int{3, 0, 1},
	},

	// 可以有多个 testcase
}

func Test_movesToStamp(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, movesToStamp(tc.stamp, tc.target), "输入:%v", tc)
	}
}

func Benchmark_movesToStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			movesToStamp(tc.stamp, tc.target)
		}
	}
}
