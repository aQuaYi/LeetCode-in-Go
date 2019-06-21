package problem1035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans int
}{

	{
		[]int{1, 4, 2},
		[]int{1, 2, 4},
		2,
	},

	{
		[]int{2, 5, 1, 2, 5},
		[]int{10, 5, 2, 1, 5, 2},
		3,
	},

	{
		[]int{1, 3, 7, 1, 7, 5},
		[]int{1, 9, 2, 5, 1},
		2,
	},

	// 可以有多个 testcase
}

func Test_maxUncrossedLines(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxUncrossedLines(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_maxUncrossedLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxUncrossedLines(tc.A, tc.B)
		}
	}
}
