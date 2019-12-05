package problem0342

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans bool
}{

	{2, false},
	{16, true},
	{5, false},
	{-5, false},

	// 可以有多个 testcase
}

func Test_isPowerOfFour(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPowerOfFour(tc.num), "输入:%v", tc)
	}
}

func Benchmark_isPowerOfFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPowerOfFour(tc.num)
		}
	}
}
