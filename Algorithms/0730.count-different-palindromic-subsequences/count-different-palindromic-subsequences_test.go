package problem0730

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
		"aacaca",
		11,
	},

	{
		"aabacabaa",
		28,
	},

	{
		"aabcdaa",
		13,
	},

	{
		"abcda",
		8,
	},

	{
		"bccb",
		6,
	},

	{
		"abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba",
		104860361,
	},

	{
		"bdddbcbaaacdcbacdbbcdcacdddacdbaddbabdbcbccdbdcadbbbbdabacdddcbacabcdcdaccdcbdaadbbbbdbddccacdadadbcbddbdccbcacbdcabdbdcccbdccbcdbadbcdbacadadaaadacdaadaabdbcccddcccdcadaccdaacbdddcbadadcadcbcbaaadcaabadbcadcdaaddacaddbadbbdbbcacaaacbbaadcaadbbaadbbbabbddcacaaabddbccbcacaadcbabcaaacdbdbcdabdbdddbcddaaabcadccbcdccdcdadbcddcdaacbaadbcbcdcbadbccbdbaabbabcadacbdcdbaadbcacdcbabcdaabdccadabccdccdcabbcabcbccdaaaddabbcdabcbdbddbbddaadbcbbddabccadccbcccdcdcbbbdaccdcbdccbdbbaacbabacdcabcdbcbadadadbcdaccaadcdcbbbcdcbcbcbaaaadcabacbdbdbcaddbadcaabdbacbaacbdcadcdbacbbbbabacdbcdaacdcbdadaaddbacccbccccaacabacddbbbbcbcbccbbccaddbaabcdbdbaaaccddacaaccdaccdcbcacbccbbabcbcadbcbdacdcbdacacadcdbbbbdcbddcddddaadbdacbcddbdacbadaacdabaddddaddbcbbbdaddcaaaabaacdbacaddcadbcaaabdbcdacaabcdcdcdddacdacbabdcaabdaaadddbcdbbdaacacacbabaadaaabcdbbabacddcdbaacaaadddddcabdabaababcbdcbadbddaccddcbabbdbcbaacdbbbdbdadddacdcddaacdbbadbacbcaccacbbbbdcadabadaccbcdcbcbbacacdaddabdddbdccbbccdaacbcbbabaddabcbbacccbbadbadbccccccc",
		115571776,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countPalindromicSubsequences(tc.s), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countPalindromicSubsequences(tc.s)
		}
	}
}
