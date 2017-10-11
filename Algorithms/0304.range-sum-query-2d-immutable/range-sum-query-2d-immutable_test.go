package Problem0304

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var matrix = [][]int{}

// tcs is testcase slice
var tcs = []struct {
	row1, col1, row2, col2 int
	ans                    int
}{

	{
		2, 1, 4, 3,
		8,
	},

	// 可以有多个 testcase
}

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	nm := Constructor(matrix)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nm.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2), "输入:%v", tc)
	}
}

func Benchmark_Constructor(b *testing.B) {
	nm := Constructor(matrix)

	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nm.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2)
		}
	}
}
