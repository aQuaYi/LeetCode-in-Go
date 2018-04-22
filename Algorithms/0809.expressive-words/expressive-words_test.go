package problem0809

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S     string
	words []string
	ans   int
}{

	{
		"heeellooo",
		[]string{"hello", "hi", "helo"},
		1,
	},

	{
		"zzzzzyyyyy",
		[]string{"zzyy", "zy", "zyy"},
		3,
	},

	// 可以有多个 testcase
}

func Test_expressiveWords(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, expressiveWords(tc.S, tc.words), "输入:%v", tc)
	}
}

func Benchmark_expressiveWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			expressiveWords(tc.S, tc.words)
		}
	}
}
