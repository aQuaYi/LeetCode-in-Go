package problem0386

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans []int
}{

	{
		13,
		[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
	},

	// 可以有多个 testcase
}

func Test_lexicalOrder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, lexicalOrder(tc.n), "输入:%v", tc)
	}
}

func Benchmark_lexicalOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lexicalOrder(tc.n)
		}
	}
}
