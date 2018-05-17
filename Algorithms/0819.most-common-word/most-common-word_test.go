package problem0819

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	paragraph string
	banned    []string
	ans       string
}{

	{
		"Bob hit a ball, the hit BALL flew far after it was hit.",
		[]string{"hit"},
		"ball",
	},

	// 可以有多个 testcase
}

func Test_mostCommonWord(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, mostCommonWord(tc.paragraph, tc.banned), "输入:%v", tc)
	}
}

func Benchmark_mostCommonWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mostCommonWord(tc.paragraph, tc.banned)
		}
	}
}
