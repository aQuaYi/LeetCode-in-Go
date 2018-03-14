package problem0500

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   []string
}{

	{
		[]string{"Hello", "Alaska", "Dad", "Peace"},
		[]string{"Alaska", "Dad"},
	},

	// 可以有多个 testcase
}

func Test_findWords(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findWords(tc.words), "输入:%v", tc)
	}
}

func Benchmark_findWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findWords(tc.words)
		}
	}
}
