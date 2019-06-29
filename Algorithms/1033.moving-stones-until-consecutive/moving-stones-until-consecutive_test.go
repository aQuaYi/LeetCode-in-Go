package problem1033

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   int
	b   int
	c   int
	ans []int
}{

	{
		1,
		2,
		5,
		[]int{1, 2},
	},

	{
		4,
		3,
		2,
		[]int{0, 0},
	},

	{
		3,
		5,
		1,
		[]int{1, 2},
	},

	// 可以有多个 testcase
}

func Test_numMovesStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numMovesStones(tc.a, tc.b, tc.c), "输入:%v", tc)
	}
}

func Benchmark_numMovesStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numMovesStones(tc.a, tc.b, tc.c)
		}
	}
}
