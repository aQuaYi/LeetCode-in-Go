package problem0743

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	times [][]int
	N     int
	K     int
	ans   int
}{

	{
		[][]int{
			{3, 5, 78},
			{2, 1, 1},
			{1, 3, 0},
			{4, 3, 59},
			{5, 3, 85},
			{5, 2, 22},
			{2, 4, 23},
			{1, 4, 43},
			{4, 5, 75},
			{5, 1, 15},
			{1, 5, 91},
			{4, 1, 16},
			{3, 2, 98},
			{3, 4, 22},
			{5, 4, 31},
			{1, 2, 0},
			{2, 5, 4},
			{4, 2, 51},
			{3, 1, 36},
			{2, 3, 59}},
		5,
		5,
		31,
	},

	{
		[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
		4,
		2,
		2,
	},

	{
		[][]int{{2, 1, 1}, {2, 3, 1}},
		4,
		2,
		-1,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, networkDelayTime(tc.times, tc.N, tc.K), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			networkDelayTime(tc.times, tc.N, tc.K)
		}
	}
}
