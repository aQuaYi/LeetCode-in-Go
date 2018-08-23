package problem0888

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	B   []int
	ans []int
}{

	{
		[]int{1, 1},
		[]int{2, 2},
		[]int{1, 2},
	},

	{
		[]int{1, 2},
		[]int{2, 3},
		[]int{1, 2},
	},

	{
		[]int{2},
		[]int{1, 3},
		[]int{2, 3},
	},

	{
		[]int{1, 2, 5},
		[]int{2, 4},
		[]int{5, 4},
	},

	// 可以有多个 testcase
}

func Test_fairCandySwap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fairCandySwap(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_fairCandySwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fairCandySwap(tc.A, tc.B)
		}
	}
}
