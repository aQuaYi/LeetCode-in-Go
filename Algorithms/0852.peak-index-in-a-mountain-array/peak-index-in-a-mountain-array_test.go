package problem0852

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8},
		8,
	},

	{
		[]int{8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		1,
	},

	{
		[]int{0, 1, 0},
		1,
	},

	{
		[]int{0, 2, 1, 0},
		1,
	},

	// 可以有多个 testcase
}

func Test_peakIndexInMountainArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, peakIndexInMountainArray(tc.A), "输入:%v", tc)
	}
}

func Benchmark_peakIndexInMountainArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			peakIndexInMountainArray(tc.A)
		}
	}
}
