package problem1007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans int
}{

	{
		[]int{2, 1, 1, 3, 2, 1, 2, 2, 1},
		[]int{3, 2, 3, 1, 3, 2, 3, 3, 2},
		-1,
	},

	{
		[]int{5, 2, 6, 2, 3, 2},
		[]int{2, 1, 2, 4, 2, 2},
		2,
	},

	{
		[]int{2, 1, 2, 4, 2, 2},
		[]int{5, 2, 6, 2, 3, 2},
		2,
	},

	{
		[]int{3, 5, 1, 2, 3},
		[]int{3, 6, 3, 3, 4},
		-1,
	},

	// 可以有多个 testcase
}

func Test_minDominoRotations(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minDominoRotations(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_minDominoRotations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDominoRotations(tc.A, tc.B)
		}
	}
}
