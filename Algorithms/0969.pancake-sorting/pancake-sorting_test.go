package problem0969

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A []int
}{

	{
		[]int{3, 2, 4, 1},
	},

	{
		[]int{1, 2, 3},
	},

	// 可以有多个 testcase
}

func Test_pancakeSort(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		A := make([]int, len(tc.A))
		copy(A, tc.A)
		ans := pancakeSort(tc.A)
		for _, k := range ans {
			swap(tc.A, k)
		}
		ast.True(sort.IntsAreSorted(tc.A), "输入 %v, 结果 %v", A, ans)
	}
}

func Benchmark_pancakeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pancakeSort(tc.A)
		}
	}
}
