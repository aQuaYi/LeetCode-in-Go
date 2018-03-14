package problem0762

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	L   int
	R   int
	ans int
}{

	{
		6,
		10,
		4,
	},

	{
		16,
		16,
		0,
	},

	{
		10,
		15,
		5,
	},

	// 可以有多个 testcase
}

func Test_countPrimeSetBits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countPrimeSetBits(tc.L, tc.R), "输入:%v", tc)
	}
}

func Benchmark_countPrimeSetBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countPrimeSetBits(tc.L, tc.R)
		}
	}
}
