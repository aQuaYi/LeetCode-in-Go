package problem0992

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{2, 1, 2, 1, 2},
		2,
		10,
	},

	{
		[]int{1, 2, 1, 2, 3},
		2,
		7,
	},

	{
		[]int{1, 2, 1, 3, 4},
		3,
		3,
	},

	// 可以有多个 testcase
}

func Test_subarraysWithKDistinct(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, subarraysWithKDistinct(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_subarraysWithKDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subarraysWithKDistinct(tc.A, tc.K)
		}
	}
}
