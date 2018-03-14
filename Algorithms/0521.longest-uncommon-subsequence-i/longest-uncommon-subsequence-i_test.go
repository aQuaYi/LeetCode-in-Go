package problem0521

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   string
	b   string
	ans int
}{

	{
		"cdc",
		"cdc",
		-1,
	},

	{
		"cdc",
		"abba",
		4,
	},

	{
		"abba",
		"cdc",
		4,
	},

	{
		"aba",
		"cdc",
		3,
	},

	// 可以有多个 testcase
}

func Test_findLUSlength(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLUSlength(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_findLUSlength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLUSlength(tc.a, tc.b)
		}
	}
}
