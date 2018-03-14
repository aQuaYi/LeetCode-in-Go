package problem0424

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	k   int
	ans int
}{

	{
		"BBBABBBCDEFGHIJKLMNOPQ",
		1,
		7,
	},

	{
		"AABABBA",
		1,
		4,
	},

	{
		"AAAA",
		2,
		4,
	},

	{
		"ABAA",
		0,
		2,
	},

	{
		"ABAB",
		2,
		4,
	},

	// 可以有多个 testcase
}

func Test_characterReplacement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, characterReplacement(tc.s, tc.k), "输入:%v", tc)
	}
}

func Benchmark_characterReplacement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			characterReplacement(tc.s, tc.k)
		}
	}
}
