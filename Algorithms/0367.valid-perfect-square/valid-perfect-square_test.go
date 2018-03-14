package problem0367

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

	{16, true},
	{64, true},
	{256, true},
	{14, false},

	// 可以有多个 testcase
}

func Test_isPerfectSquare(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPerfectSquare(tc.num), "输入:%v", tc)
	}
}

func Benchmark_isPerfectSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPerfectSquare(tc.num)
		}
	}
}
