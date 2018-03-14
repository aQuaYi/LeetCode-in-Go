package problem0436

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals []Interval
	ans       []int
}{

	{
		kit.Intss2IntervalSlice([][]int{
			{3, 4},
			{2, 3},
			{1, 2},
		}),
		[]int{-1, 0, 1},
	},

	{
		kit.Intss2IntervalSlice([][]int{
			{1, 2},
		}),
		[]int{-1},
	},

	{
		kit.Intss2IntervalSlice([][]int{
			{1, 4},
			{2, 3},
			{3, 4},
		}),
		[]int{-1, 2, -1},
	},

	// 可以有多个 testcase
}

func Test_findRightInterval(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRightInterval(tc.intervals), "输入:%v", tc)
	}
}

func Benchmark_findRightInterval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRightInterval(tc.intervals)
		}
	}
}
