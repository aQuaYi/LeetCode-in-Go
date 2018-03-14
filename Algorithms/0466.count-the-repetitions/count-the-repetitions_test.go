package problem0466

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s1  string
	n1  int
	s2  string
	n2  int
	ans int
}{

	{
		"abc",
		1,
		"abc",
		1,
		1,
	},

	{
		"ba",
		2,
		"bababa",
		1,
		0,
	},

	{
		"dcba",
		5,
		"eabcd",
		1,
		0,
	},

	{
		"dcba",
		5,
		"abcde",
		1,
		0,
	},

	{
		"aaa",
		3,
		"aa",
		1,
		4,
	},

	{
		"phqghumeaylnlfdxfircvscxggbwkfnqduxwfnfozvsrtkjprepggxrpnrvystmwcysyycqpevikeffmznimkkasvwsrenzkycxf",
		1000000,
		"xtlsgypsfadpooefxzbcoejuvpvaboygpoeylfpbnpljvrvipyamyehwqnqrqpmxujjloovaowuxwhmsncbxcoksfzkvatxdknly",
		100,
		303,
	},

	{
		"phqghumeaylnlfdxfircvscxggbwkfnqduxwfnfozvsrtkjprepggxrpnrvystmwcysyycqpevikeffmznimkkasvwsrenzkycxf",
		100,
		"xtlsgypsfa",
		1,
		49,
	},

	{
		"lovelive",
		100000,
		"lovelive",
		100000,
		1,
	},

	{
		"aaa",
		20,
		"aaaaa",
		1,
		12,
	},

	{
		"aaaaa",
		8,
		"ab",
		2,
		0,
	},

	{
		"aaaaa",
		8,
		"aa",
		2,
		10,
	},

	{
		"acb",
		6,
		"abc",
		1,
		3,
	},

	{
		"acb",
		6,
		"ab",
		2,
		3,
	},

	{
		"acb",
		4,
		"ab",
		2,
		2,
	},

	// 可以有多个 testcase
}

func Test_getMaxRepetitions(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getMaxRepetitions(tc.s1, tc.n1, tc.s2, tc.n2), "输入:%v", tc)
	}
}

func Benchmark_getMaxRepetitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getMaxRepetitions(tc.s1, tc.n1, tc.s2, tc.n2)
		}
	}
}
