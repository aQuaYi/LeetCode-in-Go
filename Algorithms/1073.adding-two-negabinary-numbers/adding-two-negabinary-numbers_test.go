package problem1073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr1 []int
	arr2 []int
	ans  []int
}{

	{
		[]int{1, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0, 0, 0, 0, 0},
		[]int{0},
	},

	{
		[]int{1},
		[]int{1, 1},
		[]int{0},
	},

	{
		[]int{1, 1, 1, 0, 1},
		[]int{1, 0, 1},
		[]int{1, 0, 1, 1, 0},
	},

	{
		[]int{1, 1, 1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 0, 0, 0, 0},
	},

	// 可以有多个 testcase
}

func Test_addNegabinary(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, addNegabinary(tc.arr1, tc.arr2), "输入:%v", tc)
	}
}

func Benchmark_addNegabinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addNegabinary(tc.arr1, tc.arr2)
		}
	}
}
