package problem1078

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text   string
	first  string
	second string
	ans    []string
}{

	{
		"alice is a good girl she is a good student",
		"a",
		"good",
		[]string{"girl", "student"},
	},

	{
		"we will we will rock you",
		"we",
		"will",
		[]string{"we", "rock"},
	},

	// 可以有多个 testcase
}

func Test_findOcurrences(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, findOcurrences(tc.text, tc.first, tc.second), "输入:%v", tc)
	}
}

func Benchmark_findOcurrences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findOcurrences(tc.text, tc.first, tc.second)
		}
	}
}
