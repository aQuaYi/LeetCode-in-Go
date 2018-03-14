package problem0766

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    bool
}{

	{
		[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
		true,
	},

	{
		[][]int{{1, 2}, {2, 2}},
		false,
	},

	// 可以有多个 testcase
}

func Test_isToeplitzMatrix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isToeplitzMatrix(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_isToeplitzMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isToeplitzMatrix(tc.matrix)
		}
	}
}
