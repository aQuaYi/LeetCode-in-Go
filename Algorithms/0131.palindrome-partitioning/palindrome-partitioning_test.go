package problem0131

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
		"cbbbcc",
		[][]string{
			[]string{"c", "b", "b", "b", "c", "c"},
			[]string{"c", "b", "b", "b", "cc"},
			[]string{"c", "b", "bb", "c", "c"},
			[]string{"c", "b", "bb", "cc"},
			[]string{"c", "bb", "b", "c", "c"},
			[]string{"c", "bb", "b", "cc"},
			[]string{"c", "bbb", "c", "c"},
			[]string{"c", "bbb", "cc"},
			[]string{"cbbbc", "c"},
		},
	},

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
