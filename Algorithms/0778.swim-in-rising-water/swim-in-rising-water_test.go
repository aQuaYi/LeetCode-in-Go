package Problem0778

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
		[][]int{{0, 2}, {1, 3}},
		3,
	},

	{
		[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}},
		16,
	},

	// 可以有多个 testcase
}

func Test_swimInWater(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, swimInWater(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_swimInWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			swimInWater(tc.grid)
		}
	}
}
