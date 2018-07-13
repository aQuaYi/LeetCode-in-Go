package problem0433

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	start string
	end   string
	bank  []string
	ans   int
}{

	{
		"AACCGGTT",
		"AAACGGTA",
		[]string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"},
		4,
	},

	{
		"AACCGGTT",
		"AACCGGTA",
		[]string{},
		-1,
	},

	{
		"AACCGGTT",
		"AACCGGTT",
		[]string{"AACCGGTT"},
		1,
	},

	{
		"AACCGGTT",
		"AACCGGTA",
		[]string{"AACCGGTA"},
		1,
	},

	{
		"AACCGGTT",
		"AAACGGTA",
		[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
		2,
	},

	{
		"AAAAACCC",
		"AACCCCCC",
		[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
		3,
	},

	// 可以有多个 testcase
}

func Test_countSegments(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minMutation(tc.start, tc.end, tc.bank), "输入:%v", tc)
	}
}

func Benchmark_countSegments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minMutation(tc.start, tc.end, tc.bank)
		}
	}
}
