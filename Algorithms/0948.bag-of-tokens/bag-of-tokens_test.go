package problem0948

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tokens []int
	P      int
	ans    int
}{

	{
		[]int{33, 4, 28, 24, 96},
		35,
		3,
	},

	{
		[]int{26},
		51,
		1,
	},

	{
		[]int{100, 200, 300, 400},
		200,
		2,
	},

	{
		[]int{100},
		50,
		0,
	},

	{
		[]int{100, 200},
		150,
		1,
	},

	// 可以有多个 testcase
}

func Test_bagOfTokensScore(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, bagOfTokensScore(tc.tokens, tc.P), "输入:%v", tc)
	}
}

func Benchmark_bagOfTokensScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			bagOfTokensScore(tc.tokens, tc.P)
		}
	}
}
