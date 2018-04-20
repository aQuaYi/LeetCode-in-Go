package problem0812

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    float64
}{

	{
		[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
		2,
	},

	// 可以有多个 testcase
}

func Test_largestTriangleArea(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largestTriangleArea(tc.points), "输入:%v", tc)
	}
}

func Benchmark_largestTriangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestTriangleArea(tc.points)
		}
	}
}
