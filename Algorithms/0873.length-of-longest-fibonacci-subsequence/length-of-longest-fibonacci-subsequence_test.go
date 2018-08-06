package problem0873

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		5,
	},

	{
		[]int{1, 3, 7, 11, 12, 14, 18},
		3,
	},

	// 可以有多个 testcase
}

func Test_lenLongestFibSubseq(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, lenLongestFibSubseq(tc.A), "输入:%v", tc)
	}
}

func Benchmark_lenLongestFibSubseq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lenLongestFibSubseq(tc.A)
		}
	}
}
