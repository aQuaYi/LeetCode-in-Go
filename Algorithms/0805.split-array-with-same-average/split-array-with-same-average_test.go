package problem0805

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans bool
}{

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 7},
		false,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		true,
	},

	// 可以有多个 testcase
}

func Test_splitArraySameAverage(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, splitArraySameAverage(tc.A), "输入:%v", tc)
	}
}

func Benchmark_splitArraySameAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			splitArraySameAverage(tc.A)
		}
	}
}
