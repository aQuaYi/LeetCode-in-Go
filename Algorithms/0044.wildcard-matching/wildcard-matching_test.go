package problem0044

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	p   string
	ans bool
}{

	{
		"aa",
		"a",
		false,
	},

	{
		"aa",
		"*",
		true,
	},

	{
		"cb",
		"?a",
		false,
	},

	{
		"adceb",
		"*a*b",
		true,
	},

	{
		"acdcb",
		"a*c?b",
		false,
	},

	// 可以有多个 testcase
}

func Test_isMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isMatch(tc.s, tc.p), "输入:%v", tc)
	}
}

func Benchmark_isMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isMatch(tc.s, tc.p)
		}
	}
}
