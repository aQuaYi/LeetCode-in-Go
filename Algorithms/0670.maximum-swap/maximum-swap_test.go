package problem0670

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans int
}{

	{
		1234,
		4231,
	},

	{
		9973,
		9973,
	},

	{
		9919,
		9991,
	},

	{
		2736,
		7236,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maximumSwap(tc.num), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximumSwap(tc.num)
		}
	}
}
