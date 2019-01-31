package problem0934

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   [][]int
	ans int
}{

	{
		[][]int{
			{0, 1, 0},
			{0, 0, 0},
			{0, 0, 1},
		},
		2,
	},

	{
		[][]int{
			{0, 1},
			{1, 0},
		},
		1,
	},

	{
		[][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1},
		},
		1,
	},

	// 可以有多个 testcase
}

func Test_shortestBridge(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, shortestBridge(tc.A), "输入:%v", tc)
	}
}

func Benchmark_shortestBridge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestBridge(tc.A)
		}
	}
}

func Test_searchForIsland_panicWithNoIsland(t *testing.T) {
	ast := assert.New(t)
	ast.Panics(func() { searchForIsland([][]int{}) })
}

func Test_shortestBridge_panicWithOnlyOneIsland(t *testing.T) {
	ast := assert.New(t)
	ast.Panics(func() { shortestBridge([][]int{{1, 0}}) })
}
