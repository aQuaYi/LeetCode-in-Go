package problem1054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	barcodes []int
}{

	{
		[]int{7, 7, 7, 8, 5, 7, 5, 5, 5, 8},
	},

	{
		[]int{2, 2, 1, 3},
	},

	{
		[]int{2, 1, 1},
	},

	{
		[]int{1, 1, 1, 2, 2, 2},
	},

	{
		[]int{1, 1, 1, 1, 2, 2, 3, 3},
	},

	// 可以有多个 testcase
}

func isDiff(A []int) bool {
	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			return false
		}
	}
	return true
}

func Test_rearrangeBarcodes(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.True(isDiff(rearrangeBarcodes(tc.barcodes)), "输入:%v", tc)
	}
}

func Benchmark_rearrangeBarcodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rearrangeBarcodes(tc.barcodes)
		}
	}
}
