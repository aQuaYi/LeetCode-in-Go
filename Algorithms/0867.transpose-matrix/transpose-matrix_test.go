package problem0867

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   [][]int
	ans [][]int
}{

	{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
	},

	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[][]int{{1, 4}, {2, 5}, {3, 6}},
	},

	// 可以有多个 testcase
}

func Test_transpose(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, transpose(tc.A), "输入:%v", tc)
	}
}

func Benchmark_transpose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			transpose(tc.A)
		}
	}
}
