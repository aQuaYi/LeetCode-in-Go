package problem1122

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr1 []int
	arr2 []int
	ans  []int
}{

	{
		[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6},
		[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
	},

	// 可以有多个 testcase
}

func Test_relativeSortArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, relativeSortArray(tc.arr1, tc.arr2), "输入:%v", tc)
	}
}

func Benchmark_relativeSortArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			relativeSortArray(tc.arr1, tc.arr2)
		}
	}
}
