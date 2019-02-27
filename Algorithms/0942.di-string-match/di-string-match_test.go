package problem0942

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S string
}{

	{
		"IDID",
	},

	{
		"III",
	},

	{
		"DDI",
	},

	// 可以有多个 testcase
}

func isCorrect(s string, a []int) bool {
	for i, b := range s {
		switch b {
		case 'I':
			if !(a[i] < a[i+1]) {
				return false
			}
		case 'D':
			if !(a[i] > a[i+1]) {
				return false
			}
		default:
			panic("s contain not only I and D")
		}
	}
	return true
}

func Test_diStringMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		a := diStringMatch(tc.S)
		ast.True(isCorrect(tc.S, a))
	}
}

func Benchmark_diStringMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			diStringMatch(tc.S)
		}
	}
}
