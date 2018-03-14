package problem0231

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans bool
}{

	{1, true},
	{2, true},
	{3, false},
	{-1, false},

	// 可以有多个 testcase
}

func Test_isPowerOfTwo(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPowerOfTwo(tc.n), "输入:%v", tc)
	}
}

func Benchmark_isPowerOfTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPowerOfTwo(tc.n)
		}
	}
}
