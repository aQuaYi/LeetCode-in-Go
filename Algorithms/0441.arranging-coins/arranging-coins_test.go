package problem0441

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{
		1804289383,
		60070,
	},

	{
		10,
		4,
	},

	{
		5,
		2,
	},

	{
		8,
		3,
	},

	// 可以有多个 testcase
}

func Test_arrangeCoins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, arrangeCoins(tc.n), "输入:%v", tc)
	}
}

func Benchmark_arrangeCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			arrangeCoins(tc.n)
		}
	}
}
