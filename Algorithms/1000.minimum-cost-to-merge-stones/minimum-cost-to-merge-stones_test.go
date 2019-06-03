package problem1000

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones []int
	K      int
	ans    int
}{

	{
		[]int{3, 5, 1, 2, 6},
		3,
		25,
	},

	{
		[]int{3, 2, 4, 1},
		2,
		20,
	},

	{
		[]int{3, 2, 4, 1},
		3,
		-1,
	},

	// 可以有多个 testcase
}

func Test_mergeStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, mergeStones(tc.stones, tc.K), "输入:%v", tc)
	}
}

func Benchmark_mergeStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mergeStones(tc.stones, tc.K)
		}
	}
}
