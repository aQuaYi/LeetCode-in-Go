package problem0929

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	emails []string
	ans    int
}{

	{
		[]string{"test.email@leetcode.com", "test.e.mai+bob.cathy@leetcode.com"},
		2,
	},

	{
		[]string{"test.email+alex@leetcode.com", "test.e.mai+bob.cathy@leetcode.com"},
		2,
	},

	{
		[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
		2,
	},

	// 可以有多个 testcase
}

func Test_numUniqueEmails(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numUniqueEmails(tc.emails), "输入:%v", tc)
	}
}

func Benchmark_numUniqueEmails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numUniqueEmails(tc.emails)
		}
	}
}
