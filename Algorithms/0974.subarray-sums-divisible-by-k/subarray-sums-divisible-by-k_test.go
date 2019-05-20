package problem0974

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{4, 5, 0, -2, -3, 1},
		5,
		7,
	},

	// 可以有多个 testcase
}

func Test_subarraysDivByK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, subarraysDivByK(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_subarraysDivByK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			subarraysDivByK(tc.A, tc.K)
		}
	}
}
