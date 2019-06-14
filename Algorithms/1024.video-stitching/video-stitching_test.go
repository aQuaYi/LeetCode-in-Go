package problem1024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	clips [][]int
	T     int
	ans   int
}{

	{
		[][]int{{5, 7}, {1, 8}, {0, 0}, {2, 3}, {4, 5}, {0, 6}, {5, 10}, {7, 10}},
		5,
		1,
	},

	{
		[][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}},
		10,
		3,
	},

	{
		[][]int{{1, 1}, {1, 2}},
		5,
		-1,
	},

	{
		[][]int{{0, 1}, {1, 2}},
		5,
		-1,
	},

	{
		[][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}},
		9,
		3,
	},

	{
		[][]int{{0, 4}, {2, 8}},
		5,
		2,
	},

	// 可以有多个 testcase
}

func Test_videoStitching(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, videoStitching(tc.clips, tc.T), "输入:%v", tc)
	}
}

func Benchmark_videoStitching(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			videoStitching(tc.clips, tc.T)
		}
	}
}
