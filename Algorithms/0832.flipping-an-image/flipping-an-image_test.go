package problem0832

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
		[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
		[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
	},

	{
		[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
		[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
	},

	// 可以有多个 testcase
}

func Test_flipAndInvertImage(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, flipAndInvertImage(tc.A), "输入:%v", tc)
	}
}

func Benchmark_flipAndInvertImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			flipAndInvertImage(tc.A)
		}
	}
}
