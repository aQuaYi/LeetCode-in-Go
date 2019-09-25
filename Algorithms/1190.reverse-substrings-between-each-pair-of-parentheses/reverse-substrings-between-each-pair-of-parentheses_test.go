package problem1190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	{
		"(x(ab)y(cd)z)",
		"zcdyabx",
	},

	{
		"(uoy(love)mi(hate)ko)",
		"okhateimloveyou",
	},

	{
		"(u(love)i(love)k)",
		"kloveiloveu",
	},

	{
		"(abcd)",
		"dcba",
	},

	{
		"(u(love)i)",
		"iloveu",
	},

	{
		"(ed(et(oc))el)",
		"leetcode",
	},

	{
		"a(bcdefghijkl(mno)p)q",
		"apmnolkjihgfedcbq",
	},

	// 可以有多个 testcase
}

func Test_reverseParentheses(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reverseParentheses(tc.s), "输入:%v", tc)
	}
}

func Benchmark_reverseParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseParentheses(tc.s)
		}
	}
}
