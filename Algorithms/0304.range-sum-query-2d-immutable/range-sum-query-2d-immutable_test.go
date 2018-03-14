package problem0304

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var matrix = [][]int{
	[]int{3, 0, 1, 4, 2},
	[]int{5, 6, 3, 2, 1},
	[]int{1, 2, 0, 1, 5},
	[]int{4, 1, 0, 1, 7},
	[]int{1, 0, 3, 0, 5},
}

// tcs is testcase slice
var tcs = []struct {
	row1, col1, row2, col2 int
	ans                    int
}{

	{
		2, 1, 4, 3,
		8,
	},

	{
		1, 1, 2, 2,
		11,
	},

	{
		1, 2, 2, 4,
		12,
	},

	// 可以有多个 testcase
}

func Test_SumRegion(t *testing.T) {
	ast := assert.New(t)

	nm := Constructor(matrix)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nm.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2), "输入:%v", tc)
	}
}

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	actual := Constructor([][]int{})
	expected := [][]int{[]int{0}}

	ast.Equal(expected, actual.dp)
}

func Benchmark_SumRegion(b *testing.B) {
	nm := Constructor(matrix)

	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nm.SumRegion(tc.row1, tc.col1, tc.row2, tc.col2)
		}
	}
}
