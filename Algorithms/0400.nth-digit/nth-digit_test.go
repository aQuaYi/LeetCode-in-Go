package problem0400

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
		2147483647,
		2,
	},

	{
		3,
		3,
	},

	{
		11,
		0,
	},

	// 可以有多个 testcase
}

func Test_findNthDigit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findNthDigit(tc.n), "输入:%v", tc)
	}
}

func Benchmark_findNthDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findNthDigit(tc.n)
		}
	}
}
