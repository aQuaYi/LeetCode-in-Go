package problem0803

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	hits [][]int
	ans  []int
}{

	{
		[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}},
		[][]int{{1, 0}},
		[]int{2},
	},

	{
		[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}},
		[][]int{{1, 1}, {1, 0}},
		[]int{0, 0},
	},

	// 可以有多个 testcase
}

func Test_hitBricks(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, hitBricks(tc.grid, tc.hits), "输入:%v", tc)
	}
}

func Benchmark_hitBricks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hitBricks(tc.grid, tc.hits)
		}
	}
}
