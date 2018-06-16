package problem0844

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	T   string
	ans bool
}{

	{
		"#ab#c",
		"#ad#c",
		true,
	},

	{
		"ab#c",
		"ad#c",
		true,
	},

	{
		"ab##",
		"c#d#",
		true,
	},

	{
		"a##c",
		"#a#c",
		true,
	},

	{
		"a#c",
		"b",
		false,
	},

	// 可以有多个 testcase
}

func Test_backspaceCompare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, backspaceCompare(tc.S, tc.T), "输入:%v", tc)
	}
}

func Benchmark_backspaceCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			backspaceCompare(tc.S, tc.T)
		}
	}
}
