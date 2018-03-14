package problem0326

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

	{-1, false},
	{8, false},
	{2, false},
	{3, true},
	{81, true},
	{243, true},

	// 可以有多个 testcase
}

func Test_isPowerOfThree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPowerOfThree(tc.n), "输入:%v", tc)
	}
}

func Benchmark_isPowerOfThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPowerOfThree(tc.n)
		}
	}
}
