package problem0691

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stickers []string
	target   string
	ans      int
}{

	{
		[]string{"thethat", "with", "example", "science"},
		"thehat",
		1,
	},

	{
		[]string{"t", "example", "science", "with"},
		"thehat",
		3,
	},

	{
		[]string{"example", "science", "with"},
		"thehat",
		3,
	},

	{
		[]string{"with", "example", "science"},
		"thehat",
		3,
	},

	{
		[]string{"notice", "possible"},
		"basicbasic",
		-1,
	},

	// 可以有多个 testcase
}

func Test_minStickers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minStickers(tc.stickers, tc.target), "输入:%v", tc)
	}
}

func Benchmark_minStickers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minStickers(tc.stickers, tc.target)
		}
	}
}
