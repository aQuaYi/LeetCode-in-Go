package problem0522

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	strs []string
	ans  int
}{

	{
		[]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"},
		6,
	},

	{
		[]string{"aaa", "aaa", "bb", "c", "c", "d", "d"},
		2,
	},

	{
		[]string{"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"},
		-1,
	},

	{
		[]string{"aabbcc", "aabbcc", "bc", "bcc", "aabbccc"},
		7,
	},

	{
		[]string{"abcd", "abcd", "abc", "bc"},
		-1,
	},

	{
		[]string{"abcd", "abcd", "abc", "cde"},
		3,
	},

	{
		[]string{"aba", "cdc", "eae"},
		3,
	},

	// 可以有多个 testcase
}

func Test_findLUSlength(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findLUSlength(tc.strs), "输入:%v", tc)
	}
}

func Benchmark_findLUSlength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLUSlength(tc.strs)
		}
	}
}
