package problem1163

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
		"xbylisvborylklftlkcioajuxwdhahdgezvyjbgaznzayfwsaumeccpfwamfzmkinezzwobllyxktqeibfoupcpptncggrdqbkji",
		"zzwobllyxktqeibfoupcpptncggrdqbkji",
	},

	{
		"abab",
		"bab",
	},

	{
		"leetcode",
		"tcode",
	},

	// 可以有多个 testcase
}

func Test_lastSubstring(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, lastSubstring(tc.s), "输入:%v", tc)
	}
}

func Benchmark_lastSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lastSubstring(tc.s)
		}
	}
}
