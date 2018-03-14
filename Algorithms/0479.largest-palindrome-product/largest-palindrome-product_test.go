package problem0479

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
		1,
		9,
	},

	{
		2,
		987,
	},

	{
		3,
		123,
	},

	{
		4,
		597,
	},

	{
		5,
		677,
	},

	{
		6,
		1218,
	},

	{
		7,
		877,
	},

	{
		8,
		475,
	},

	// 可以有多个 testcase
}

func Test_largestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largestPalindrome(tc.n), "输入:%v", tc)
	}
}

func Benchmark_largestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestPalindrome(tc.n)
		}
	}
}
