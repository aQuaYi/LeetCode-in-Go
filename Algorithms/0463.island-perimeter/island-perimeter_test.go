package problem0463

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
		[][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		},
		16,
	},

	// 可以有多个 testcase
}

func Test_islandPerimeter(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, islandPerimeter(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_islandPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			islandPerimeter(tc.grid)
		}
	}
}
