package problem0868

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
		41,
		3,
	},

	{
		22,
		2,
	},

	{
		5,
		2,
	},

	{
		6,
		1,
	},

	{
		8,
		0,
	},

	// 可以有多个 testcase
}

func Test_binaryGap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, binaryGap(tc.N), "输入:%v", tc)
	}
}

func Benchmark_binaryGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			binaryGap(tc.N)
		}
	}
}
