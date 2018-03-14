package problem0592

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	ans        string
}{

	{
		"-1/2+1/2",
		"0/1",
	},

	{
		"-1/2+1/2+1/3",
		"1/3",
	},

	{
		"1/3-1/2",
		"-1/6",
	},

	{
		"5/3+1/3",
		"2/1",
	},

	// 可以有多个 testcase
}

func Test_fractionAddition(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fractionAddition(tc.expression), "输入:%v", tc)
	}
}

func Benchmark_fractionAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fractionAddition(tc.expression)
		}
	}
}
