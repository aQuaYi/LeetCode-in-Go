package problem1079

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tiles string
	ans   int
}{

	{
		"ABCDEFG",
		13699,
	},
	{
		"AAB",
		8,
	},

	{
		"AAABBC",
		188,
	},

	// 可以有多个 testcase
}

func Test_numTilePossibilities(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numTilePossibilities(tc.tiles), "输入:%v", tc)
	}
}

func Benchmark_numTilePossibilities(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numTilePossibilities(tc.tiles)
		}
	}
}
