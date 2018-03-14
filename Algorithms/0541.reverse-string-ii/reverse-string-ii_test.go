package problem0541

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	k   int
	ans string
}{

	{
		"abcdefg",
		3,
		"cbadefg",
	},

	{
		"abcdefg",
		2,
		"bacdfeg",
	},

	// 可以有多个 testcase
}

func Test_reverseStr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reverseStr(tc.s, tc.k), "输入:%v", tc)
	}
}

func Benchmark_reverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseStr(tc.s, tc.k)
		}
	}
}
