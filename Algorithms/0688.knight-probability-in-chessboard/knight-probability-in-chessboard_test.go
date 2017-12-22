package Problem0688

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
		3, 2, 0, 0,
		0.0625,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, knightProbability(tc.n, tc.k, tc.r, tc.c), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			knightProbability(tc.n, tc.k, tc.r, tc.c)
		}
	}
}
