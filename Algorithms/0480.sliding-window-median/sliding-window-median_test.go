package problem0480

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  []float64
}{

	{
		[]int{1, 2},
		1,
		[]float64{1, 2},
	},

	{
		[]int{1, 2},
		2,
		[]float64{1.5},
	},

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]float64{1, -1, -1, 3, 5, 6},
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		14,
		[]float64{7.5},
	},

	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		3,
		[]float64{8, 7, 6, 5, 4, 3, 2},
	},

	// 可以有多个 testcase
}

func Test_medianSlidingWindow(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, medianSlidingWindow(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_medianSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			medianSlidingWindow(tc.nums, tc.k)
		}
	}
}

// 这个是为了覆盖率
func Test_pop(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	k := 14
	w := newWindow(nums, k)
	w.min.Pop()
	w.max.Pop()
}
