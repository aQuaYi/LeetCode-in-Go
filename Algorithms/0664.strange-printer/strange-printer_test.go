package problem0664

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"ccdcadbddbaddcbccdcdabcbcddbccdcbad",
		17,
	},

	{
		"tbgtgb",
		4,
	},

	{
		"aaabbb",
		2,
	},

	{
		"",
		0,
	},

	{
		"abababa",
		4,
	},

	{
		"aba",
		2,
	},

	// 可以有多个 testcase
}

func Test_strangePrinter(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, strangePrinter(tc.s), "输入:%v", tc)
	}
}

func Benchmark_strangePrinter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			strangePrinter(tc.s)
		}
	}
}
