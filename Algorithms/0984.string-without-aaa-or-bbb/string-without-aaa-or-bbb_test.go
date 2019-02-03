package problem0984

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   int
	B   int
	ans string
}{

	{
		1,
		2,
		"abb",
	},

	{
		4,
		1,
		"aabaa",
	},

	// 可以有多个 testcase
}

func isCorrect(A, B int, s string) bool {
	if A+B != len(s) ||
		len(strings.Replace(s, "a", "", -1)) != B ||
		len(strings.Replace(s, "b", "", -1)) != A {
		return false
	}
	if strings.Contains(s, "aaa") ||
		strings.Contains(s, "bbb") {
		return false
	}
	return true
}

func Test_strWithout3a3b(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.True(isCorrect(tc.A, tc.B, strWithout3a3b(tc.A, tc.B)))
	}
}

func Benchmark_strWithout3a3b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			strWithout3a3b(tc.A, tc.B)
		}
	}
}
