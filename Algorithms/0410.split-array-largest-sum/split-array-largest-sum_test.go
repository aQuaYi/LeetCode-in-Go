package problem0410

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	m    int
	ans  int
}{

	{
		[]int{1, 3, 5, 7, 9, 11},
		3,
		16,
	},

	{
		[]int{1, 2, 2, 2},
		3,
		3,
	},

	{
		[]int{1, 2, 2, 2},
		1,
		7,
	},

	{
		[]int{7, 2, 5, 10, 8},
		2,
		18,
	},

	{
		[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
		8,
		25,
	},

	// 可以有多个 testcase
}

func Test_splitArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, splitArray(tc.nums, tc.m), "输入:%v", tc)
	}
}

func Benchmark_splitArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			splitArray(tc.nums, tc.m)
		}
	}
}
