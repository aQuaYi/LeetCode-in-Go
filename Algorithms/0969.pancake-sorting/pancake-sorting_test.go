package problem0969

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans []int
}{

	{
		[]int{3, 2, 4, 1},
		[]int{4, 2, 4, 3},
	},

	{
		[]int{1, 2, 3},
		[]int{},
	},

	// 可以有多个 testcase
}

func Test_pancakeSort(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, pancakeSort(tc.A), "输入:%v", tc)
	}
}

func Benchmark_pancakeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pancakeSort(tc.A)
		}
	}
}
