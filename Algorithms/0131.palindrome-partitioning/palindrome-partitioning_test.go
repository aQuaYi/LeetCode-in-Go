package Problem0131

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans [][]string
}{

	{
		"efe",
		[][]string{
			[]string{"e", "f", "e"},
			[]string{"efe"},
		},
	},

	{
		"a",
		[][]string{
			[]string{"a"},
		},
	},

	{
		"aab",
		[][]string{
			[]string{"a", "a", "b"},
			[]string{"aa", "b"},
		},
	},

	// 可以有多个 testcase
}

func Test_partition(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, partition(tc.s), "输入:%v", tc)
	}
}

func Benchmark_partition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			partition(tc.s)
		}
	}
}
