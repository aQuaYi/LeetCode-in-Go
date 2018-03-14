package problem0212

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	words []string
	ans   []string
}{

	{
		[][]byte{
			[]byte("a"),
		},
		[]string{"a", "a"},
		[]string{"a"},
	},

	{
		[][]byte{},
		[]string{"a", "a"},
		nil,
	},

	{
		[][]byte{
			[]byte{},
		},
		[]string{"a", "a"},
		nil,
	},

	{
		[][]byte{
			[]byte("oaan"),
			[]byte("etae"),
			[]byte("ihkr"),
			[]byte("iflv"),
		},
		[]string{"oath", "pea", "eat", "rain", "eakat"},
		[]string{"oath", "eat"},
	},

	// 可以有多个 testcase
}

func Test_findWords(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findWords(tc.board, tc.words), "输入:%v", tc)
	}
}

func Benchmark_findWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findWords(tc.board, tc.words)
		}
	}
}
