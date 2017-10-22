package Problem0329

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    int
}{

	{[][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	},
		4},

	{[][]int{
		[]int{3, 4, 5},
		[]int{3, 2, 6},
		[]int{2, 2, 1},
	},
		4},

	// 可以有多个 testcase
}

func Test_longestIncreasingPath(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestIncreasingPath(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_longestIncreasingPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestIncreasingPath(tc.matrix)
		}
	}
}
