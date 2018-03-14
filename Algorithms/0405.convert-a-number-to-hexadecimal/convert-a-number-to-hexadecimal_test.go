package problem0405

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

	{
		0,
		"0",
	},

	{
		26,
		"1a",
	},

	{
		-1,
		"ffffffff",
	},

	// 可以有多个 testcase
}

func Test_toHex(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, toHex(tc.num), "输入:%v", tc)
	}
}

func Benchmark_toHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			toHex(tc.num)
		}
	}
}
