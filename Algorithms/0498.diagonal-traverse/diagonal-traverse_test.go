package problem0498

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    []int
}{

	{
		[][]int{},
		[]int{},
	},

	{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
	},

	// 可以有多个 testcase
}

func Test_findDiagonalOrder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findDiagonalOrder(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_findDiagonalOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findDiagonalOrder(tc.matrix)
		}
	}
}
