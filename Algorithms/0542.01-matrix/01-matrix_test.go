package problem0542

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    [][]int
}{

	{
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	},

	{
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		},
	},

	// 可以有多个 testcase
}

func Test_updateMatrix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, updateMatrix(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_updateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			updateMatrix(tc.matrix)
		}
	}
}
