package problem0187

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans []string
}{

	{
		"",
		nil,
	},

	{
		"ACACA",
		nil,
	},

	{
		"ACACACACACACA",
		[]string{"ACACACACAC", "CACACACACA"},
	},

	{
		"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		[]string{"AAAAACCCCC", "CCCCCAAAAA"},
	},

	// 可以有多个 testcase
}

func Test_findRepeatedDnaSequences(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRepeatedDnaSequences(tc.s), "输入:%v", tc)
	}
}

func Benchmark_findRepeatedDnaSequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRepeatedDnaSequences(tc.s)
		}
	}
}
