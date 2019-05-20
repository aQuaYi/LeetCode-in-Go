package problem0975

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{10, 13, 12, 14, 15},
		2,
	},

	{
		[]int{2, 3, 1, 1, 4},
		3,
	},

	{
		[]int{5, 1, 3, 4, 2},
		3,
	},

	// 可以有多个 testcase
}

func Test_oddEvenJumps(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, oddEvenJumps(tc.A), "输入:%v", tc)
	}
}

func Benchmark_oddEvenJumps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			oddEvenJumps(tc.A)
		}
	}
}
