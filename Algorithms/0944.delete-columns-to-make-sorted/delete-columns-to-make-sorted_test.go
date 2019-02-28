package problem0944

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans int
}{

	{
		[]string{"cba", "daf", "ghi"},
		1,
	},

	{
		[]string{"a", "b"},
		0,
	},

	{
		[]string{"zyx", "wvu", "tsr"},
		3,
	},

	// 可以有多个 testcase
}

func Test_minDeletionSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minDeletionSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_minDeletionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDeletionSize(tc.A)
		}
	}
}
