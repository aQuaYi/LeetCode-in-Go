package problem0421

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
		[]int{0},
		0,
	},

	{
		[]int{3, 10, 5, 25, 2, 8},
		28,
	},

	// 可以有多个 testcase
}

func Test_findMaximumXOR(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMaximumXOR(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_findMaximumXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMaximumXOR(tc.nums)
		}
	}
}
