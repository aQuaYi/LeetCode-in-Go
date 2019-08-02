package problem0090

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans [][]int
}{

	{
		[]int{1, 2, 2},
		[][]int{[]int{}, []int{1}, []int{1, 2}, []int{1, 2, 2}, []int{2}, []int{2, 2}},
	},

	{
		[]int{2, 1, 2, 1, 3},
		[][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}},

	},

	// 可以有多个 testcase
}

func Test_subsetsWithDup(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, subsetsWithDup(tc.A), "输入:%v", tc)
	}
}

func Benchmark_subsetsWithDup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subsetsWithDup(tc.A)
		}
	}
}
