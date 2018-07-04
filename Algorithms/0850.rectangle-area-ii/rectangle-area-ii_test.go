package problem0850

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rectangles [][]int
	ans        int
}{

	{
		[][]int{{0, 0, 2, 2}, {3, 0, 5, 2}},
		8,
	},

	{
		[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}},
		6,
	},

	{
		[][]int{{0, 0, 1000000000, 1000000000}},
		49,
	},

	{
		[][]int{{0, 0, 1000000000, 1000000000}, {0, 1000000000, 1000000000, 2000000000}},
		98,
	},

	// 可以有多个 testcase
}

func Test_rectangleArea(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, rectangleArea(tc.rectangles), "输入:%v", tc)
	}
}

func Benchmark_rectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rectangleArea(tc.rectangles)
		}
	}
}
