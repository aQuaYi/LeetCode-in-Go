package problem0205

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans bool
}{

	{"ba", "a", false},
	{"a", "a", true},
	{"aba", "baa", false},
	{"ab", "aa", false},
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},

	// 可以有多个 testcase
}

func Test_isIsomorphic(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isIsomorphic(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_isIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isIsomorphic(tc.s, tc.t)
		}
	}
}
