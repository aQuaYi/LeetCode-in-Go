package problem0989

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans []int
}{

	{
		[]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3},
		516,
		[]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9},
	},

	{
		[]int{1, 2, 0, 0},
		34,
		[]int{1, 2, 3, 4},
	},

	{
		[]int{2, 7, 4},
		181,
		[]int{4, 5, 5},
	},

	{
		[]int{2, 1, 5},
		806,
		[]int{1, 0, 2, 1},
	},

	{
		[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		1,
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},

	// 可以有多个 testcase
}

func Test_addToArrayForm(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, addToArrayForm(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_addToArrayForm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addToArrayForm(tc.A, tc.K)
		}
	}
}
