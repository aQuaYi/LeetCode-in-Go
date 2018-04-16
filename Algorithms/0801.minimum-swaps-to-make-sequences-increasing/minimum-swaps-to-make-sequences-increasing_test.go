package problem0801

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans int
}{

	{
		[]int{2, 3, 2, 5, 6},
		[]int{0, 1, 4, 4, 5},
		1,
	},

	{
		[]int{0, 2, 5, 8, 9, 10, 12, 14, 18, 19, 20, 20, 24, 27, 28, 31, 33, 34, 36, 38},
		[]int{1, 2, 5, 7, 8, 9, 11, 17, 15, 16, 19, 21, 28, 29, 30, 31, 33, 34, 38, 39},
		2,
	},

	{
		[]int{0, 3, 4, 9, 10},
		[]int{2, 3, 7, 5, 6},
		1,
	},

	{
		[]int{3, 3, 8, 9, 10},
		[]int{1, 7, 4, 6, 8},
		1,
	},

	{
		[]int{1, 3, 5, 4},
		[]int{1, 2, 3, 7},
		1,
	},

	// 可以有多个 testcase

}

func Test_minSwap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minSwap(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_minSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minSwap(tc.A, tc.B)
		}
	}
}
