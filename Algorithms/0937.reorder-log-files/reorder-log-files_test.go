package problem0937

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	logs []string
	ans  []string
}{

	{
		[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
		[]string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
	},

	// 可以有多个 testcase
}

func Test_reorderLogFiles(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, reorderLogFiles(tc.logs), "输入:%v", tc)
	}
}

func Benchmark_reorderLogFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderLogFiles(tc.logs)
		}
	}
}
