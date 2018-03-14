package problem0264

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

	{
		11,
		15,
	},

	{
		12,
		16,
	},

	{
		1690,
		2123366400,
	},

	{
		1,
		1,
	},

	// 可以有多个 testcase
}

func Test_nthUglyNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nthUglyNumber(tc.n), "输入:%v", tc)
	}
}

func Benchmark_nthUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nthUglyNumber(tc.n)
		}
	}
}
