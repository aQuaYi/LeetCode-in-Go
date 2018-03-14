package problem0692

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	k     int
	ans   []string
}{

	{
		[]string{"i", "love", "leetcode", "i", "love", "coding"},
		2,
		[]string{"i", "love"},
	},

	{
		[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
		4,
		[]string{"the", "is", "sunny", "day"},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, topKFrequent(tc.words, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			topKFrequent(tc.words, tc.k)
		}
	}
}
