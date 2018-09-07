package problem0892

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{{1, 0}, {0, 2}},
		16,
	},

	{
		[][]int{{2}},
		10,
	},

	{
		[][]int{{1, 2}, {3, 4}},
		34,
	},

	{
		[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		32,
	},

	{
		[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
		46,
	},

	// 可以有多个 testcase
}

func Test_surfaceArea(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, surfaceArea(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_surfaceArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			surfaceArea(tc.grid)
		}
	}
}
