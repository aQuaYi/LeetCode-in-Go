package problem0930

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	S   int
	ans int
}{

	{
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		6,
		3,
	},

	{
		[]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		0,
		27,
	},

	{
		[]int{0, 0, 0, 0, 1},
		2,
		0,
	},

	{
		[]int{0, 0, 0, 0, 0},
		0,
		15,
	},

	{
		[]int{1, 0, 1, 0, 1},
		2,
		4,
	},

	// 可以有多个 testcase
}

func Test_numSubarraysWithSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numSubarraysWithSum(tc.A, tc.S), "输入:%v", tc)
	}
}

func Benchmark_numSubarraysWithSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSubarraysWithSum(tc.A, tc.S)
		}
	}
}
