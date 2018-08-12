package problem0860

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	bills []int
	ans   bool
}{

	{
		[]int{5, 5, 5, 10, 20},
		true,
	},

	{
		[]int{5, 5, 5, 20, 10, 20},
		false,
	},

	{
		[]int{5, 5, 10},
		true,
	},

	{
		[]int{10, 10},
		false,
	},

	{
		[]int{5, 5, 10, 10, 20},
		false,
	},

	// 可以有多个 testcase
}

func Test_lemonadeChange(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, lemonadeChange(tc.bills), "输入:%v", tc)
	}
}

func Benchmark_lemonadeChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lemonadeChange(tc.bills)
		}
	}
}
