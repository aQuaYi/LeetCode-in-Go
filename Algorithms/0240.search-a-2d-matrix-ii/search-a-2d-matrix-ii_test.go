package problem0240

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = [][]int{
	[]int{1, 4, 7, 11, 15},
	[]int{2, 5, 8, 12, 19},
	[]int{3, 6, 9, 16, 22},
	[]int{10, 13, 14, 17, 24},
	[]int{18, 21, 23, 26, 30},
}

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	target int
	ans    bool
}{

	{
		m,
		5,
		true,
	},

	{
		m,
		15,
		true,
	},
	{
		m,
		20,
		false,
	},

	{
		[][]int{},
		5,
		false,
	},

	{
		[][]int{[]int{}},
		5,
		false,
	},
	// 可以有多个 testcase
}

func Test_searchMatrix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, searchMatrix(tc.matrix, tc.target), "输入:%v", tc)
	}
}

func Benchmark_searchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			searchMatrix(tc.matrix, tc.target)
		}
	}
}
