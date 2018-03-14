package problem0395

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	k   int
	ans int
}{

	{
		"abcdedghijklmnopqrstuvwxyz",
		2,
		0,
	},

	{
		"aaabb",
		4,
		0,
	},

	{
		"acabb",
		2,
		2,
	},

	{
		"aaabb",
		3,
		3,
	},

	{
		"ababbc",
		2,
		5,
	},

	// 可以有多个 testcase
}

func Test_longestSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, longestSubstring(tc.s, tc.k), "输入:%v", tc)
	}
}

func Benchmark_longestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestSubstring(tc.s, tc.k)
		}
	}
}
