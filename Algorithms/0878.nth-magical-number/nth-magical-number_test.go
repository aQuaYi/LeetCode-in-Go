package problem0878

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	A   int
	B   int
	ans int
}{

	{
		1000000000,
		40000,
		40000,
		999720007,
	},

	{
		1,
		2,
		3,
		2,
	},

	{
		4,
		2,
		3,
		6,
	},

	{
		5,
		2,
		4,
		10,
	},

	{
		3,
		6,
		4,
		8,
	},

	// 可以有多个 testcase
}

func Test_nthMagicalNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nthMagicalNumber(tc.N, tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_nthMagicalNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nthMagicalNumber(tc.N, tc.A, tc.B)
		}
	}
}
