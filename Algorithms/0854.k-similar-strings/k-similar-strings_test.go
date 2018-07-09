package problem0854

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans int
}{

	{
		"ab",
		"ab",
		0,
	},

	{
		"ab",
		"ba",
		1,
	},

	{
		"abc",
		"bca",
		2,
	},

	{
		"abac",
		"baca",
		2,
	},

	{
		"aabc",
		"abca",
		2,
	},

	// 可以有多个 testcase
}

func Test_kSimilarity(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kSimilarity(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_kSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kSimilarity(tc.A, tc.B)
		}
	}
}
