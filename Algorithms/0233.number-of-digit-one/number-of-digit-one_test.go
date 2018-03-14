package problem0233

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
		13,
		6,
	},

	{
		111,
		36,
	},

	{
		123456,
		93553,
	},

	{
		12345678,
		11824417,
	},

	// 可以有多个 testcase
}

func Test_countDigitOne(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countDigitOne(tc.n), "输入:%v", tc)
	}
}

func Benchmark_countDigitOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countDigitOne(tc.n)
		}
	}
}
