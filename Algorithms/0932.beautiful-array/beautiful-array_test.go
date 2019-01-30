package problem0932

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans []int
}{

	{
		4,
		[]int{2, 1, 4, 3},
	},

	{
		5,
		[]int{3, 1, 2, 5, 4},
	},

	// 可以有多个 testcase
}

func Test_beautifulArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, beautifulArray(tc.N), "输入:%v", tc)
	}
}

func Benchmark_beautifulArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			beautifulArray(tc.N)
		}
	}
}
