package problem0943

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A        []string
	expected string
}{

	{
		[]string{"fgbabcde", "abcdefgb", "bcdei"},
		"abcdefgbabcdei",
	},

	{
		[]string{"cedefifgstkyxfcuajfa", "ooncedefifgstkyxfcua", "assqjfwarvjcjedqtoz", "fcuajfassqjfwarvjc", "fwarvjcjedqtozctcd", "zppedxfumcfsngp", "kyxfcuajfassqjfwa", "fumcfsngphjyfhhwkqa", "fassqjfwarvjcjedq", "ppedxfumcfsngphjyf", "dqtozctcdk"},
		"ooncedefifgstkyxfcuajfassqjfwarvjcjedqtozctcdkzppedxfumcfsngphjyfhhwkqa",
	},
	{
		[]string{"alex", "loves", "leetcode"},
		"alexlovesleetcode",
	},

	{
		[]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"},
		"gctaagttcatgcatc",
	},

	// 可以有多个 testcase
}

func isCorrect(A []string, expected, actual string) bool {
	if expected == actual {
		return true
	}

	if len(expected) != len(actual) {
		return false
	}

	for _, sub := range A {
		if !strings.Contains(actual, sub) {
			return false
		}
	}

	return true
}

func Test_shortestSuperstring(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		actual := shortestSuperstring(tc.A)
		ast.True(isCorrect(tc.A, tc.expected, actual), "输入:%v", tc)
	}
}

func Benchmark_shortestSuperstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestSuperstring(tc.A)
		}
	}
}
