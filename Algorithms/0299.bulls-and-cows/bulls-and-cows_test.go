package problem0299

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	secret string
	guess  string
	ans    string
}{

	{
		"1807",
		"7810",
		"1A3B",
	},

	{
		"1123456789",
		"0111111111",
		"1A1B",
	},

	// 可以有多个 testcase
}

func Test_getHint(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getHint(tc.secret, tc.guess), "输入:%v", tc)
	}
}

func Benchmark_getHint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getHint(tc.secret, tc.guess)
		}
	}
}
