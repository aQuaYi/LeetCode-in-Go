package problem0520

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	word string
	ans  bool
}{

	{
		"leetcode",
		true,
	},

	{
		"USA",
		true,
	},

	{
		"aG",
		false,
	},

	{
		"FlaG",
		false,
	},

	// 可以有多个 testcase
}

func Test_detectCapitalUse(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, detectCapitalUse(tc.word), "输入:%v", tc)
	}
}

func Benchmark_detectCapitalUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			detectCapitalUse(tc.word)
		}
	}
}
