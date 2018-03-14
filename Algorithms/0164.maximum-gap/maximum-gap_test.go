package problem0164

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1},
		0,
	},

	{
		[]int{1, 10000000},
		9999999,
	},

	{
		[]int{9, 8, 7, 6, 5, 4, 3, 3, 1},
		2,
	},

	// 可以有多个 testcase
}

func Test_maximumGap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maximumGap(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_maximumGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximumGap(tc.nums)
		}
	}
}
