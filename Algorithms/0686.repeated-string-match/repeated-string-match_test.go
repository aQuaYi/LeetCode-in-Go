package problem0686

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a, b string
	ans  int
}{

	{
		"ba",
		"ab",
		2,
	},

	{
		"abc",
		"aabcbabcc",
		-1,
	},

	{
		"abcd",
		"cdabdab",
		-1,
	},

	{
		"aa",
		"b",
		-1,
	},

	{
		"aa",
		"a",
		1,
	},

	{
		"abcd",
		"bcabcdbc",
		-1,
	},

	{
		"abcd",
		"abcdb",
		-1,
	},

	{
		"abababaaba",
		"aabaaba",
		2,
	},

	{
		"abcd",
		"cdabcdab",
		3,
	},

	{
		"bb",
		"bbbbbbb",
		4,
	},

	{
		"a",
		"aa",
		2,
	},

	{
		"abc",
		"cabcabca",
		4,
	},
	// 可以有多个 testcase
}

func Test_repeatedStringMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, repeatedStringMatch(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			repeatedStringMatch(tc.a, tc.b)
		}
	}
}
