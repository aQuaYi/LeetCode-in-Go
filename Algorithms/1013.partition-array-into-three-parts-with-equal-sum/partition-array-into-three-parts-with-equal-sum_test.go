package problem1013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{6, 1, 1, 13, -1, 0, -10, 20},
		false,
	},

	{
		[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
		true,
	},

	{
		[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1, 2},
		false,
	},

	{
		[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
		true,
	},

	{
		[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
		false,
	},

	// 可以有多个 testcase
}

func Test_canThreePartsEqualSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, canThreePartsEqualSum(tc.A), "输入:%v", tc)
	}
}

func Benchmark_canThreePartsEqualSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canThreePartsEqualSum(tc.A)
		}
	}
}
