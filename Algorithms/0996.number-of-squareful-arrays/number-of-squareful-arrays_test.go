package problem0996

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
		[]int{7, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		12,
	},

	{
		[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		1,
	},

	{
		[]int{1, 17, 8},
		2,
	},

	{
		[]int{2, 2, 2},
		1,
	},

	// 可以有多个 testcase
}

func Test_numSquarefulPerms(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, numSquarefulPerms(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numSquarefulPerms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSquarefulPerms(tc.A)
		}
	}
}
