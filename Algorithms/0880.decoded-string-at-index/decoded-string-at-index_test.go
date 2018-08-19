package problem0880

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	K   int
	ans string
}{

	{
		"leet2code3",
		10,
		"o",
	},

	{
		"ha22",
		5,
		"h",
	},

	{
		"a2345678999999999999999",
		1,
		"a",
	},

	// 可以有多个 testcase
}

func Test_decodeAtIndex(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, decodeAtIndex(tc.S, tc.K), "输入:%v", tc)
	}
}

func Benchmark_decodeAtIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			decodeAtIndex(tc.S, tc.K)
		}
	}
}
