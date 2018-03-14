package problem0452

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    int
}{

	{
		[][]int{},
		0,
	},

	{
		[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		2,
	},

	// 可以有多个 testcase
}

func Test_findMinArrowShots(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMinArrowShots(tc.points), "输入:%v", tc)
	}
}

func Benchmark_findMinArrowShots(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMinArrowShots(tc.points)
		}
	}
}
