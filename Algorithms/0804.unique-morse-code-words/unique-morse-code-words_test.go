package problem0804

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   int
}{

	{
		[]string{"gin", "zen", "gig", "msg"},
		2,
	},

	{
		[]string{
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyz",
		},
		1,
	},

	// 可以有多个 testcase
}

func Test_uniqueMorseRepresentations(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, uniqueMorseRepresentations(tc.words), "输入:%v", tc)
	}
}

func Benchmark_uniqueMorseRepresentations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniqueMorseRepresentations(tc.words)
		}
	}
}
