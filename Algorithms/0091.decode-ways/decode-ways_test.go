package problem0091

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
		"203",
		1,
	},

	{
		"546313216586746123648561321321564321657651231365456475642312356486765123156465132546543213546316543465465465465427",
		5038848,
	},

	{
		"027",
		0,
	},

	{
		"27",
		1,
	},

	{
		"1",
		1,
	},

	{
		"119",
		3,
	},

	{
		"2011",
		2,
	},

	{
		"1192",
		3,
	},

	{
		"1102",
		1,
	},

	{
		"1128",
		3,
	},

	{
		"1",
		1,
	},

	{
		"12",
		2,
	},

	{
		"121",
		3,
	},

	{
		"1212",
		5,
	},

	{
		"12121",
		8,
	},

	{
		"121212",
		13,
	},

	{
		"1212121",
		21,
	},

	{
		"12121212",
		34,
	},

	{
		"12",
		2,
	},

	{
		"226",
		3,
	},

	// 可以有多个 testcase
}

func Test_numDecodings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numDecodings(tc.s), "输入:%v", tc)
	}
}

func Benchmark_numDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numDecodings(tc.s)
		}
	}
}
