package problem0273

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
	{50868, "Fifty Thousand Eight Hundred Sixty Eight"},
	{100, "One Hundred"},
	{30, "Thirty"},
	{21, "Twenty One"},
	{0, "Zero"},
	{123, "One Hundred Twenty Three"},
	{12345, "Twelve Thousand Three Hundred Forty Five"},
	{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	{101, "One Hundred One"},
	{1, "One"},
	{11, "Eleven"},
	{111, "One Hundred Eleven"},
	{1111, "One Thousand One Hundred Eleven"},
	{11111, "Eleven Thousand One Hundred Eleven"},
	{111111, "One Hundred Eleven Thousand One Hundred Eleven"},
	{1111111, "One Million One Hundred Eleven Thousand One Hundred Eleven"},
	{11111111, "Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
	{111111111, "One Hundred Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
	{1111111111, "One Billion One Hundred Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
	{1000000000, "One Billion"},

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
