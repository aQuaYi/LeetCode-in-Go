package problem0960

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans int
}{

	{
		[]string{"aaaabaa"},
		1,
	},

	{
		[]string{"babca", "bbazb"},
		3,
	},

	{
		[]string{"edcba"},
		4,
	},

	{
		[]string{"ghi", "def", "abc"},
		0,
	},

	{
		[]string{
			"deehdecfcgegffegghfhfaagcaaffbfahcfaghgdfggbbddbff",
			"dchhgcbahdbdgbbaafhbgfggcbebfacdebhfgcfaafcgbgbggg",
			"hehdggagfabdfbhegebhaaddcaghhegeegdgegagehhdhheecd",
			"fhbbagbdffedafacbeahddbgagggdafceeabaffhhhhcedcfbh",
			"caaefgdgefeahcgfgccaacdfabdgdbdhdbhbhfadbeaegbbdce",
			"habgbahaeebeacccbdhfhddegfebheeffdbcbgfahhgbhcheeb",
			"gfaaedgcachcehgdghebbhegbfagdgcdcgebddbdccbedbbhcd",
			"badaebdbdgeadbfgchaegaddgdhdgaeaaedagacgbdecfdghca",
			"eefcceffcggeefbecaceadbdcfccgbfgffgahfgeebafdcfhhb",
			"aeebdahabfecbafgbcchbhdeecaadaccbahghcbdcfaeagehha",
		},
		48,
	},

	// 可以有多个 testcase
}

func Test_minDeletionSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minDeletionSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minDeletionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDeletionSize(tc.A)
		}
	}
}
