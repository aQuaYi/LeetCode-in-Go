package problem0816

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []string
}{

	{
		"(123)",
		[]string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"},
	},

	{
		"(00011)",
		[]string{"(0.001, 1)", "(0, 0.011)"},
	},

	{
		"(0123)",
		[]string{"(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"},
	},

	{
		"(100)",
		[]string{"(10, 0)"},
	},

	// 可以有多个 testcase
}

func Test_ambiguousCoordinates(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := ambiguousCoordinates(tc.S)
		sort.Strings(ans)
		sort.Strings(tc.ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_ambiguousCoordinates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			ambiguousCoordinates(tc.S)
		}
	}
}
