package problem0719

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  int
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		36,
		8,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		8,
		1,
	},

	{
		[]int{1, 3, 1},
		1,
		0,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, smallestDistancePair(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestDistancePair(tc.nums, tc.k)
		}
	}
}
