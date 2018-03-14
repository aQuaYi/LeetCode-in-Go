package problem0640

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	equation string
	ans      string
}{

	{
		"-x=-1",
		"x=1",
	},

	{
		"x+5-3+x=6+x-2",
		"x=2",
	},

	{
		"x=x",
		"Infinite solutions",
	},

	{
		"2x=x",
		"x=0",
	},

	{
		"2x+3x-6x=x+2",
		"x=-1",
	},

	{
		"x=x+2",
		"No solution",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, solveEquation(tc.equation), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solveEquation(tc.equation)
		}
	}
}
