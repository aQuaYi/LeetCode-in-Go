package problem0343

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{2, 1},
	{3, 2},
	{58, 1549681956},
	{9, 27},
	{10, 36},
	{11, 54},

	// 可以有多个 testcase
}

func Test_integerBreak(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, integerBreak(tc.n), "输入:%v", tc)
	}
}

func Benchmark_integerBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			integerBreak(tc.n)
		}
	}
}
