package problem0357

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
	{1, 10},
	{2, 91},
	{3, 739},
	{4, 5275},
	{5, 32491},
	{6, 168571},
	{7, 712891},
	{8, 2345851},
	{9, 5611771},
	{10, 8877691},
	{11, 8877691},

	// 可以有多个 testcase
}

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countNumbersWithUniqueDigits(tc.n), "输入:%v", tc)
	}
}

func Benchmark_countNumbersWithUniqueDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countNumbersWithUniqueDigits(tc.n)
		}
	}
}
