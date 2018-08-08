package problem0709

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	str string
	ans string
}{

	{
		"Hello",
		"hello",
	},

	{
		"here",
		"here",
	},

	{
		"LOVELY",
		"lovely",
	},

	// 可以有多个 testcase
}

func Test_toLowerCase(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, toLowerCase(tc.str), "输入:%v", tc)
	}
}

func Benchmark_toLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			toLowerCase(tc.str)
		}
	}
}
