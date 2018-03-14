package problem0172

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
		100000,
		24999,
	},

	{
		10000,
		2499,
	},

	{
		1000,
		249,
	},

	{
		100,
		24,
	},

	{
		25,
		6,
	},

	{
		10,
		2,
	},

	{
		3,
		0,
	},

	// 可以有多个 testcase
}

func Test_trailingZeroes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, trailingZeroes(tc.n), "输入:%v", tc)
	}
}

func Benchmark_trailingZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trailingZeroes(tc.n)
		}
	}
}
