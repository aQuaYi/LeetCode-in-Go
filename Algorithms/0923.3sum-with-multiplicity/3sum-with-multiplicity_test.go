package problem0923

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A      []int
	target int
	ans    int
}{

	{
		[]int{6, 6, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		8,
		22,
	},

	{
		[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		8,
		20,
	},

	{
		[]int{2, 2, 2, 2},
		6,
		4,
	},

	{
		[]int{1, 1, 2, 2, 2, 2},
		5,
		12,
	},

	// 可以有多个 testcase
}

func Test_threeSumMulti(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, threeSumMulti(tc.A, tc.target), "输入:tc.A, tc.target", tc)
	}
}

func Benchmark_threeSumMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			threeSumMulti(tc.A, tc.target)
		}
	}
}
