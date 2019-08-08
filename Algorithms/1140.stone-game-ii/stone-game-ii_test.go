package problem1140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	piles []int
	ans   int
}{

	{
		[]int{2, 7, 9, 4, 4},
		10,
	},

	// 可以有多个 testcase
}

func Test_stoneGameII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, stoneGameII(tc.piles), "输入:%v", tc)
	}
}

func Benchmark_stoneGameII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			stoneGameII(tc.piles)
		}
	}
}
