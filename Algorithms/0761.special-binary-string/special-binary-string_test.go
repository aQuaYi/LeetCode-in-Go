package problem0761

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans string
}{

	{
		"101100",
		"110010",
	},

	{
		"11011000",
		"11100100",
	},

	// 可以有多个 testcase
}

func Test_makeLargestSpecial(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, makeLargestSpecial(tc.S), "输入:%v", tc)
	}
}

func Benchmark_makeLargestSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			makeLargestSpecial(tc.S)
		}
	}
}
