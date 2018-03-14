package problem0556

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
		12443322,
		13222344,
	},

	{
		2321,
		3122,
	},

	{
		2147483467,
		2147483476,
	},

	{
		2147483647,
		-1,
	},

	{
		12,
		21,
	},

	{
		21,
		-1,
	},

	// 可以有多个 testcase
}

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreaterElement(tc.n), "输入:%v", tc)
	}
}

func Benchmark_nextGreaterElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextGreaterElement(tc.n)
		}
	}
}
