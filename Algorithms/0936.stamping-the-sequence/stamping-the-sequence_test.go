package problem0936

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stamp  string
	target string
}{

	{
		"ab",
		"aba",
	},

	{
		"abc",
		"ababc",
	},

	{
		"abca",
		"aabcaca",
	},

	//
}

func isCorrect(stamp, target string, sequence []int) bool {
	if len(sequence) > 10*len(target) {
		return false
	}
	sb := []byte(stamp)
	lt := len(target)
	tb := make([]byte, lt*2)
	for _, i := range sequence {
		copy(tb[i:], sb)
	}
	stb := string(tb[:lt])
	if stb != target && len(sequence) != 0 {
		fmt.Println(stamp, target, sequence)
		return false
	}
	return true
}

func Test_isCorrect(t *testing.T) {
	ast := assert.New(t)
	ast.True(isCorrect("aba", "ababa", []int{0, 2}))
	ast.True(isCorrect("aba", "ababa", []int{2, 0}))
}

func Test_movesToStamp(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		seq := movesToStamp(tc.stamp, tc.target)
		ast.True(isCorrect(tc.stamp, tc.target, seq))
	}
}

func Benchmark_movesToStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			movesToStamp(tc.stamp, tc.target)
		}
	}
}
