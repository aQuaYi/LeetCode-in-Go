package problem0345

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

	{"hello", "holle"},

	{"leetcode", "leotcede"},

	{"aA", "Aa"},

	// 可以有多个 testcase
}

func Test_reverseVowels(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reverseVowels(tc.s), "输入:%v", tc)
	}
}

func Benchmark_reverseVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseVowels(tc.s)
		}
	}
}
