package problem0415

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num1 string
	num2 string
	ans  string
}{

	{
		"12",
		"3",
		"15",
	},

	{
		"2",
		"3",
		"5",
	},

	{
		"7",
		"8",
		"15",
	},

	// 可以有多个 testcase
}

func Test_addStrings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, addStrings(tc.num1, tc.num2), "输入:%v", tc)
	}
}

func Benchmark_addStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addStrings(tc.num1, tc.num2)
		}
	}
}
