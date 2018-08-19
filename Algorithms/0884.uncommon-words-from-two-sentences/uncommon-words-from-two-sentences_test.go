package problem0884

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans []string
}{

	{
		"this apple is sweet",
		"this apple is sour",
		[]string{"sweet", "sour"},
	},

	{
		"apple apple",
		"banana",
		[]string{"banana"},
	},

	// 可以有多个 testcase
}

func Test_uncommonFromSentences(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		sort.Strings(tc.ans)
		actual := uncommonFromSentences(tc.A, tc.B)
		sort.Strings(actual)
		ast.Equal(tc.ans, actual, "输入:%v", tc)
	}
}

func Benchmark_uncommonFromSentences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uncommonFromSentences(tc.A, tc.B)
		}
	}
}
