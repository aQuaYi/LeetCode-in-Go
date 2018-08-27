package problem0893

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []string
	ans int
}{

	{
		[]string{"a", "b", "c", "a", "c", "c"},
		3,
	},

	{
		[]string{"aa", "bb", "ab", "ba"},
		4,
	},

	{
		[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
		3,
	},

	{
		[]string{"abcd", "cdab", "adcb", "cbad"},
		1,
	},

	// 可以有多个 testcase
}

func Test_numSpecialEquivGroups(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSpecialEquivGroups(tc.A), "输入:%v", tc)
	}
}

func Benchmark_numSpecialEquivGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSpecialEquivGroups(tc.A)
		}
	}
}
