package problem0829

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		21,
		4,
	},

	{
		1,
		1,
	},

	{
		3,
		2,
	},

	{
		1000000000,
		10,
	},

	{
		51,
		4,
	},

	{
		9,
		3,
	},

	{
		5,
		2,
	},

	{
		15,
		4,
	},

	// 可以有多个 testcase
}

func Test_consecutiveNumbersSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, consecutiveNumbersSum(tc.N), "输入:%v", tc)
	}
}

func Benchmark_consecutiveNumbersSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			consecutiveNumbersSum(tc.N)
		}
	}
}
