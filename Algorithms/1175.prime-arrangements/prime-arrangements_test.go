package problem1175

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{
		1,
		1,
	},

	{
		5,
		12,
	},

	{
		100,
		682289015,
	},

	// 可以有多个 testcase
}

func Test_numPrimeArrangements(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, numPrimeArrangements(tc.n), "输入:%v", tc)
	}
}

func Benchmark_numPrimeArrangements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numPrimeArrangements(tc.n)
		}
	}
}
