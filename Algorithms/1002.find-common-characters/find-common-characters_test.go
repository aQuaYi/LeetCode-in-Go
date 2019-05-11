package problem1002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans []string
}{

	{
		[]string{"bella", "label", "roller"},
		[]string{"e", "l", "l"},
	},

	{
		[]string{"cool", "lock", "cook"},
		[]string{"c", "o"},
	},

	// 可以有多个 testcase
}

func Test_commonChars(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, commonChars(tc.A), "输入:%v", tc)
	}
}

func Benchmark_commonChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			commonChars(tc.A)
		}
	}
}
