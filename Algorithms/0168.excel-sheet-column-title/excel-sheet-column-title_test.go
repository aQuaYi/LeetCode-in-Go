package problem0168

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans string
}{

	{
		100,
		"CV",
	},

	{
		1000,
		"ALL",
	},

	{
		53,
		"BA",
	},

	{
		52,
		"AZ",
	},

	{
		26,
		"Z",
	},

	{
		1,
		"A",
	},

	{
		28,
		"AB",
	},

	// 可以有多个 testcase
}

func Test_convertToTitle(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, convertToTitle(tc.n), "输入:%v", tc)
	}
}

func Benchmark_convertToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			convertToTitle(tc.n)
		}
	}
}
