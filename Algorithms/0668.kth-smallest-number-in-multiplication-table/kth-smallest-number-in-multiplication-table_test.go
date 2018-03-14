package problem0668

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	m   int
	n   int
	k   int
	ans int
}{

	{
		11,
		13,
		57,
		24,
	},

	{
		30000,
		30000,
		900000000,
		900000000,
	},

	{
		30000,
		3,
		90000,
		90000,
	},

	{
		300,
		300,
		90000,
		90000,
	},

	{
		3000,
		3000,
		9000000,
		9000000,
	},

	{
		3,
		8,
		19,
		14,
	},

	{
		3,
		6,
		13,
		9,
	},

	{
		3,
		3,
		5,
		3,
	},

	{
		2,
		3,
		6,
		6,
	},

	// 可以有多个 testcase
}

func Test_findKthNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthNumber(tc.m, tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthNumber(tc.m, tc.n, tc.k)
		}
	}
}
