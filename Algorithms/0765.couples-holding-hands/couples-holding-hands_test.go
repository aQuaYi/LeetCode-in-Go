package problem0765

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	row []int
	ans int
}{

	{
		[]int{1, 4, 0, 5, 8, 7, 6, 3, 2, 9},
		3,
	},

	{
		[]int{0, 2, 1, 3},
		1,
	},

	{
		[]int{3, 2, 0, 1},
		0,
	},

	{
		[]int{3, 4, 2, 5, 0, 1},
		1,
	},

	{
		[]int{3, 4, 2, 0, 5, 1},
		2,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 0},
		3,
	},

	{
		[]int{0, 5, 3, 1, 4, 2},
		2,
	},

	// 可以有多个 testcase
}

func Test_minSwapsCouples(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minSwapsCouples(tc.row), "输入:%v", tc)
	}
}

func Benchmark_minSwapsCouples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minSwapsCouples(tc.row)
		}
	}
}

func Benchmark_delet2AtOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minSwapsCouples([]int{0, 2, 1, 3, 4, 6, 5, 7, 8, 10, 9, 11})
	}
}
