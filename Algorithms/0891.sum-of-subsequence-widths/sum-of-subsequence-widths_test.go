package problem0891

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
		[]int{2, 1, 3},
		6,
	},

	// 可以有多个 testcase
}

func Test_sumSubseqWidths(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, sumSubseqWidths(tc.A), "输入:%v", tc)
	}
}

func Benchmark_sumSubseqWidths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumSubseqWidths(tc.A)
		}
	}
}
