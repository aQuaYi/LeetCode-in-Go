package problem0313

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n      int
	primes []int
	ans    int
}{

	{
		1,
		[]int{2, 7, 13, 19},
		1,
	},

	{
		12,
		[]int{2, 7, 13, 19},
		32,
	},

	// 可以有多个 testcase
}

func Test_nthSuperUglyNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nthSuperUglyNumber(tc.n, tc.primes), "输入:%v", tc)
	}
}

func Benchmark_nthSuperUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nthSuperUglyNumber(tc.n, tc.primes)
		}
	}
}
