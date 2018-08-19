package problem0880

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
		"a2345678999999999999999",
		12314145213412,
		"a",
	},

	{
		"a2b3c4d5e6f7g8h9",
		3,
		"b",
	},

	{
		"leet2code3for2you2",
		160,
		"y",
	},

	{
		"leet2code3for2you2",
		89,
		"t",
	},

	{
		"leet2code3for2you2",
		13,
		"l",
	},

	{
		"leet2code3for2you2",
		90,
		"c",
	},

	{
		"leet2code3",
		10,
		"o",
	},

	{
		"leet2code3for2you",
		16,
		"t",
	},

	{
		"leet2code3for2you2",
		85,
		"t",
	},

	{
		"ha22",
		5,
		"h",
	},

	{
		"a2345678999999999999999",
		1,
		"a",
	},

	// 可以有多个 testcase
}

func Test_decodeAtIndex(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, decodeAtIndex(tc.S, tc.K), "输入:%v", tc)
	}
}

func Test_decodeAtIndex_2(t *testing.T) {
	ast := assert.New(t)

	s := "leat2code3for2you2"
	ansStr := "leatleatcodeleatleatcodeleatleatcodeforleatleatcodeleatleatcodeleatleatcodeforyouleatleatcodeleatleatcodeleatleatcodeforleatleatcodeleatleatcodeleatleatcodeforyou"
	for i := range ansStr {
		expectd := ansStr[i : i+1]
		actual := decodeAtIndex(s, i+1)
		msg := fmt.Sprintf("%s 的第 %d 个字符应该是 %s", s, i+1, expectd)
		ast.Equal(expectd, actual, msg)
	}
}

func Benchmark_decodeAtIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			decodeAtIndex(tc.S, tc.K)
		}
	}
}
