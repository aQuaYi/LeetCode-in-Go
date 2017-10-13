package Problem0301

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans []string
}{
	{"()())()", []string{"()()()", "(())()"}},
	{"(a)())()", []string{"(a)()()", "(a())()"}},
	{")(", []string{""}},

	// 可以有多个 testcase
}

func Test_removeInvalidParentheses(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeInvalidParentheses(tc.s), "输入:%v", tc)
	}
}

func Benchmark_removeInvalidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeInvalidParentheses(tc.s)
		}
	}
}
