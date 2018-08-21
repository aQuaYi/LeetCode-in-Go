package problem0885

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	R   int
	C   int
	r0  int
	c0  int
	ans [][]int
}{

	{
		1,
		4,
		0,
		0,
		[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	},

	{
		5,
		6,
		1,
		4,
		[][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}},
	},

	// 可以有多个 testcase
}

func Test_spiralMatrixIII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, spiralMatrixIII(tc.R, tc.C, tc.r0, tc.c0), "输入:%v", tc)
	}
}

func Benchmark_spiralMatrixIII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			spiralMatrixIII(tc.R, tc.C, tc.r0, tc.c0)
		}
	}
}
