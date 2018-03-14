package problem0375

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

	{12, 21},
	{5, 6},
	{1, 0},
	{2, 1},
	{3, 2},
	{5, 6},
	{6, 8},
	{7, 10},
	{8, 12},
	{9, 14},
	{10, 16},
	{12, 21},
	{20, 49},
	{100, 400},
	{300, 1640},

	// 可以有多个 testcase
}

func Test_getMoneyAmount(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getMoneyAmount(tc.n), "输入:%v", tc)
	}
}

func Benchmark_getMoneyAmount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getMoneyAmount(tc.n)
		}
	}
}
