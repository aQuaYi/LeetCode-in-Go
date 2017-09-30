package Problem0198

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
		[]int{},
		0,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		20,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		25,
	},

	// 可以有多个 testcase
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, rob(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rob(tc.nums)
		}
	}
}
