package problem0704

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    int
}{

	{
		[]int{-1, 0, 3, 5, 9, 12},
		9,
		4,
	},

	{
		[]int{-1, 0, 3, 5, 9, 12},
		2,
		-1,
	},

	// 可以有多个 testcase
}

func Test_search(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, search(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			search(tc.nums, tc.target)
		}
	}
}
