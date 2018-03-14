package problem0482

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	K   int
	ans string
}{

	{
		"---",
		3,
		"",
	},

	{
		"24A0r7-4k",
		3,
		"24-A0R-74K",
	},

	{
		"2-4A0r7-4k",
		4,
		"24A0-R74K",
	},

	{
		"2-4A0r7-4k",
		3,
		"24-A0R-74K",
	},

	// 可以有多个 testcase
}

func Test_licenseKeyFormatting(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, licenseKeyFormatting(tc.S, tc.K), "输入:%v", tc)
	}
}

func Benchmark_licenseKeyFormatting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			licenseKeyFormatting(tc.S, tc.K)
		}
	}
}
