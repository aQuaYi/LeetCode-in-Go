package problem0553

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

	{[]int{1000, 100, 10, 2}, "1000/(100/10/2)"},

	{[]int{2}, "2"},

	{[]int{2, 1}, "2/1"},

	{[]int{5, 2, 1}, "5/(2/1)"},

	// 可以有多个 testcase
}

func Test_optimalDivision(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, optimalDivision(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_optimalDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			optimalDivision(tc.nums)
		}
	}
}
