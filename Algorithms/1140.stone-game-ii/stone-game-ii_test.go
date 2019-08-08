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
		[]int{8270, 7145, 575, 5156, 5126, 2905, 8793, 7817, 5532, 5726, 7071, 7730, 5200, 5369, 5763, 7148, 8287, 9449, 7567, 4850, 1385, 2135, 1737, 9511, 8065, 7063, 8023, 7729, 7084, 8407},
		98008,
	},

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
