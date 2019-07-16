package problem1090

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	values     []int
	labels     []int
	num_wanted int
	use_limit  int
	ans        int
}{

	{
		[]int{5, 4, 3, 2, 1},
		[]int{1, 1, 2, 2, 3},
		3,
		1,
		9,
	},

	{
		[]int{5, 4, 3, 2, 1},
		[]int{1, 3, 3, 3, 2},
		3,
		2,
		12,
	},

	{
		[]int{9, 8, 8, 7, 6},
		[]int{0, 0, 0, 1, 1},
		3,
		1,
		16,
	},

	{
		[]int{9, 8, 8, 7, 6},
		[]int{0, 0, 0, 1, 1},
		3,
		2,
		24,
	},

	// 可以有多个 testcase
}

func Test_largestValsFromLabels(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, largestValsFromLabels(tc.values, tc.labels, tc.num_wanted, tc.use_limit), "输入:%v", tc)
	}
}

func Benchmark_largestValsFromLabels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestValsFromLabels(tc.values, tc.labels, tc.num_wanted, tc.use_limit)
		}
	}
}
