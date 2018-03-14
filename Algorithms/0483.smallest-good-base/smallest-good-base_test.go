package problem0483

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   string
	ans string
}{

	{
		"2251799813685247",
		"2",
	},

	{
		"4681",
		"8",
	},

	{
		"16035713712910627",
		"502",
	},

	{
		"14919921443713777",
		"496",
	},

	{
		"13",
		"3",
	},

	{
		"1000000000000000000",
		"999999999999999999",
	},

	// 可以有多个 testcase
}

func Test_smallestGoodBase(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, smallestGoodBase(tc.n), "输入:%v", tc)
	}
}

func Benchmark_smallestGoodBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestGoodBase(tc.n)
		}
	}
}
