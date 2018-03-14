package problem0392

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans bool
}{

	{
		"abc",
		"ahbgdc",
		true,
	},

	{
		"axc",
		"ahbgdc",
		false,
	},

	// 可以有多个 testcase
}

func Test_isSubsequence(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isSubsequence(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_isSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSubsequence(tc.s, tc.t)
		}
	}
}
