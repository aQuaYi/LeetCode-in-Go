package problem0796

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans bool
}{

	{
		"abcde",
		"cdeab",
		true,
	},

	{
		"aa",
		"a",
		false,
	},

	{
		"abcde",
		"abced",
		false,
	},

	// 可以有多个 testcase
}

func Test_rotateString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, rotateString(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_rotateString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotateString(tc.A, tc.B)
		}
	}
}
