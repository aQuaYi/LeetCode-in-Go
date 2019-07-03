package problem1047

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans string
}{

	{
		"abbaca",
		"ca",
	},

	// 可以有多个 testcase
}

func Test_removeDuplicates(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, removeDuplicates(tc.S), "输入:%v", tc)
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicates(tc.S)
		}
	}
}
