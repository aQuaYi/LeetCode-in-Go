package Problem0273

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans string
}{
	{123, "One Hundred Twenty Three"},
	{12345, "Twelve Thousand Three Hundred Forty Five"},
	{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},

	// 可以有多个 testcase
}

func Test_numberToWords(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numberToWords(tc.num), "输入:%v", tc)
	}
}

func Benchmark_numberToWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numberToWords(tc.num)
		}
	}
}
