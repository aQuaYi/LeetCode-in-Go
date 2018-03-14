package problem0435

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals [][]int
	ans       int
}{

	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		1,
	},

	{
		[][]int{{1, 2}, {1, 2}, {1, 2}},
		2,
	},

	{
		[][]int{{2, 3}},
		0,
	},

	// 可以有多个 testcase
}

func Test_eraseOverlapIntervals(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, eraseOverlapIntervals(kit.Intss2IntervalSlice(tc.intervals)), "输入:%v", tc)
	}
}

func Benchmark_eraseOverlapIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			eraseOverlapIntervals(kit.Intss2IntervalSlice(tc.intervals))
		}
	}
}
