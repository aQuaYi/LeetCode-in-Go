package problem0564

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
		"9",
		"8",
	},

	{
		"99999999999999999",
		"100000000000000001",
	},

	{
		"8888",
		"8778",
	},

	{
		"100000000000000001",
		"99999999999999999",
	},

	{
		"77777",
		"77677",
	},

	{
		"9999",
		"10001",
	},

	{
		"121",
		"111",
	},

	{
		"2",
		"1",
	},

	{
		"10",
		"9",
	},

	{
		"1",
		"0",
	},

	{
		"4321",
		"4334",
	},

	{
		"123",
		"121",
	},

	// 可以有多个 testcase
}

func Test_nearestPalindromic(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nearestPalindromic(tc.n), "输入:%v", tc)
	}
}

func Benchmark_nearestPalindromic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nearestPalindromic(tc.n)
		}
	}
}
