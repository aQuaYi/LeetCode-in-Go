package problem0807

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
		[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}},
		35,
	},

	// 可以有多个 testcase
}

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxIncreaseKeepingSkyline(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_maxIncreaseKeepingSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxIncreaseKeepingSkyline(tc.grid)
		}
	}
}
