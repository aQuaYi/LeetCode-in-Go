package problem0459

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans bool
}{

	{
		"bb",
		true,
	},

	{
		"abab",
		true,
	},

	{
		"",
		false,
	},

	{
		"aba",
		false,
	},

	{
		"abcabcabcabc",
		true,
	},

	// 可以有多个 testcase
}

func Test_repeatedSubstringPattern(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, repeatedSubstringPattern(tc.s), "输入:%v", tc)
	}
}

func Benchmark_repeatedSubstringPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			repeatedSubstringPattern(tc.s)
		}
	}
}
