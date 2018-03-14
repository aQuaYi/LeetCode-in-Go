package problem0179

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  string
}{

	{
		[]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247},
		"9609938824824769735703560743981399",
	},

	{
		[]int{0, 0},
		"0",
	},

	{
		[]int{1, 20},
		"201",
	},

	{
		[]int{3, 30, 34, 5, 9},
		"9534330",
	},

	{
		[]int{3, 30, 34, 5, 9, 30},
		"953433030",
	},

	// 可以有多个 testcase
}

func Test_largestNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largestNumber(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_largestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestNumber(tc.nums)
		}
	}
}
