package problem0811

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	cpdomains []string
	ans       []string
}{

	{
		[]string{"9001 discuss.leetcode.com"},
		[]string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
	},

	{
		[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
	},

	// 可以有多个 testcase
}

func Test_subdomainVisits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		geted := subdomainVisits(tc.cpdomains)
		sort.Strings(geted)
		sort.Strings(tc.ans)

		ast.Equal(tc.ans, geted, "输入:%v", tc)
	}
}

func Benchmark_subdomainVisits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subdomainVisits(tc.cpdomains)
		}
	}
}
