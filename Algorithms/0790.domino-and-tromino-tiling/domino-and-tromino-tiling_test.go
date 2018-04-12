package problem0790

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		5,
		24,
	},

	{
		4,
		11,
	},

	{
		3,
		5,
	},

	{
		2,
		2,
	},

	{
		1,
		1,
	},

	{
		1000,
		979232805,
	},

	// 可以有多个 testcase
}

func Test_numTilings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numTilings(tc.N), "输入:%v", tc)
	}
}

func Benchmark_numTilings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numTilings(tc.N)
		}
	}
}
