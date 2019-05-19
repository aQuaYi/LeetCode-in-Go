package problem0912

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  []int
}{

	{
		[]int{5, 2, 3, 1},
		[]int{1, 2, 3, 5},
	},

	{
		[]int{5, 1, 1, 2, 0, 0},
		[]int{0, 0, 1, 1, 2, 5},
	},

	// 可以有多个 testcase
}

func Test_sortArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, sortArray(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_sortArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortArray(tc.nums)
		}
	}
}
