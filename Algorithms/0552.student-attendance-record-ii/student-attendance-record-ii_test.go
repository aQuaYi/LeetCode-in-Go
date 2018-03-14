package problem0552

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

	{0, 1},

	{1, 3},

	{3, 19},

	{2, 8},

	{4, 43},

	{5, 94},

	{10, 3536},

	{100, 985598218},

	{1000, 250434094},

	{10000, 67802379},

	{100000, 749184020},

	// 可以有多个 testcase
}

func Test_checkRecord(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkRecord(tc.n), "输入:%v", tc)
	}
}

func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkRecord(tc.n)
		}
	}
}
