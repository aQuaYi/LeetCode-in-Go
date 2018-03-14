package problem0150

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tokens []string
	ans    int
}{

	{
		[]string{"2", "1", "-"},
		1,
	},

	{
		[]string{"2", "1", "+", "3", "*"},
		9,
	},

	{
		[]string{"4", "13", "5", "/", "+"},
		6,
	},

	// 可以有多个 testcase
}

func Test_evalRPN(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, evalRPN(tc.tokens), "输入:%v", tc)
	}
}

func Benchmark_evalRPN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			evalRPN(tc.tokens)
		}
	}
}
