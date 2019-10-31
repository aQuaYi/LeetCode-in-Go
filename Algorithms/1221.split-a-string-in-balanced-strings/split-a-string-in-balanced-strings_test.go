package problem1221

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"RLRRLLRLRL",
		4,
	},

	{
		"RLLLLRRRLR",
		3,
	},

	{
		"LLLLRRRR",
		1,
	},

	// 可以有多个 testcase
}

func Test_balancedStringSplit(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, balancedStringSplit(tc.s), "输入:%v", tc)
	}
}

func Benchmark_balancedStringSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			balancedStringSplit(tc.s)
		}
	}
}
