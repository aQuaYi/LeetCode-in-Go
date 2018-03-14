package problem0263

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
	{
		-2147483648,
		false,
	},

	{
		0,
		false,
	},

	{
		2123366402,
		false,
	},

	{
		23,
		false,
	},

	{
		12,
		true,
	},

	{
		1,
		true,
	},

	{
		2123366400,
		true,
	},
	// 可以有多个 testcase
}

func Test_isUgly(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isUgly(tc.num), "输入:%v", tc)
	}
}

func Benchmark_isUgly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isUgly(tc.num)
		}
	}
}
