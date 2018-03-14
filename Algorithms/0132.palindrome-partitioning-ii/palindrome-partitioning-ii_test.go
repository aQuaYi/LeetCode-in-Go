package problem0132

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
		"apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp",
		452,
	},

	{
		"abcdd",
		3,
	},

	{
		"aaabaa",
		1,
	},

	{
		"aaaaacbbbc",
		1,
	},

	{
		"aab",
		1,
	},

	{
		"",
		0,
	},

	// 可以有多个 testcase
}

func Test_minCut(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		fmt.Println(len(tc.s))

		ast.Equal(tc.ans, minCut(tc.s), "输入:%v", tc)
	}
}

func Benchmark_minCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minCut(tc.s)
		}
	}
}
