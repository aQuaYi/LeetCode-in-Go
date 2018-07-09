package problem0856

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"()",
		1,
	},

	{
		"(())",
		2,
	},

	{
		"()()",
		2,
	},

	{
		"(()(()))",
		6,
	},

	// 可以有多个 testcase
}

func Test_scoreOfParentheses(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, scoreOfParentheses(tc.S), "输入:%v", tc)
	}
}

func Benchmark_scoreOfParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			scoreOfParentheses(tc.S)
		}
	}
}
