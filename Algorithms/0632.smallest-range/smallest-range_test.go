package problem0632

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums [][]int
	ans  []int
}{

	{
		[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}},
		[]int{20, 24},
	},

	{
		[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
		[]int{1, 1},
	},

	// 可以有多个 testcase
}

func Test_smallestRange(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, smallestRange(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_smallestRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestRange(tc.nums)
		}
	}
}
