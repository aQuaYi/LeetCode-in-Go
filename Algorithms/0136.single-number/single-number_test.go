package problem0136

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
		[]int{1, 1, 2},
		2,
	},

	// 可以有多个 testcase
}

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, singleNumber(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNumber(tc.nums)
		}
	}
}
