package problem0524

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	d   []string
	ans string
}{

	{
		"abpcplea",
		[]string{"ale", "apple", "monkey", "plea"},
		"apple",
	},

	{
		"a",
		[]string{"bc", "c"},
		"",
	},

	{
		"a",
		[]string{"a", "bc", "c"},
		"a",
	},

	{
		"abpcplea",
		[]string{"a", "b", "c"},
		"a",
	},

	// 可以有多个 testcase
}

func Test_findLongestWord(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLongestWord(tc.s, tc.d), "输入:%v", tc)
	}
}

func Benchmark_findLongestWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLongestWord(tc.s, tc.d)
		}
	}
}
