package problem0629

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	k   int
	ans int
}{

	{
		1000,
		500,
		955735232,
	},

	{
		10,
		3,
		155,
	},

	{
		4,
		2,
		5,
	},

	{
		3,
		0,
		1,
	},

	{
		3,
		1,
		2,
	},

	// 可以有多个 testcase
}

func Test_kInversePairs(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kInversePairs(tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_kInversePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kInversePairs(tc.n, tc.k)
		}
	}
}
