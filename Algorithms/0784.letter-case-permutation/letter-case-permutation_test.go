package problem0784

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []string
}{

	{
		"A12",
		[]string{"a12", "A12"},
	},

	{
		"a1b2",
		[]string{"a1b2", "a1B2", "A1b2", "A1B2"},
	},

	{
		"3z4",
		[]string{"3z4", "3Z4"},
	},

	{
		"12345",
		[]string{"12345"},
	},

	// 可以有多个 testcase
}

func Test_letterCasePermutation(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, letterCasePermutation(tc.S), "输入:%v", tc)
	}
}

func Benchmark_letterCasePermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			letterCasePermutation(tc.S)
		}
	}
}
