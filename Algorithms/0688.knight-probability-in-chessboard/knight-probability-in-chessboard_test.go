package problem0688

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n, k, r, c int
	ans        float64
}{

	{
		8, 30, 6, 4,
		0.00019,
	},

	{
		3, 1, 1, 1,
		0.0,
	},

	{
		3, 2, 0, 0,
		0.0625,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.InDelta(tc.ans, knightProbability(tc.n, tc.k, tc.r, tc.c), 0.000001, "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			knightProbability(tc.n, tc.k, tc.r, tc.c)
		}
	}
}
