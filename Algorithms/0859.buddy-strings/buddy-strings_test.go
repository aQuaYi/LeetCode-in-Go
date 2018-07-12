package problem0859

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans bool
}{

	{
		"ab",
		"ba",
		true,
	},

	{
		"ab",
		"ab",
		false,
	},

	{
		"aa",
		"aa",
		true,
	},

	{
		"aaaaaaabc",
		"aaaaaaacb",
		true,
	},

	{
		"",
		"aa",
		false,
	},

	// 可以有多个 testcase
}

func Test_buddyStrings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, buddyStrings(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_buddyStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			buddyStrings(tc.A, tc.B)
		}
	}
}
