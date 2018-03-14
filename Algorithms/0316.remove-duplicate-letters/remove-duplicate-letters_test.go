package problem0316

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{
	{"bbbacacca", "bac"},
	{"ccacbaba", "acb"},
	{"cbacdcbc", "acdb"},
	{"cbac", "bac"},
	{"bcabc", "abc"},

	// 可以有多个 testcase
}

func Test_removeDuplicateLetters(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeDuplicateLetters(tc.s), "输入:%v", tc)
	}
}

func Benchmark_removeDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicateLetters(tc.s)
		}
	}
}
