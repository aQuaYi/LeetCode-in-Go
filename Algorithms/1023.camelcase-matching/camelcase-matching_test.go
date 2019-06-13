package problem1023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	queries []string
	pattern string
	ans     []bool
}{

	{
		[]string{"IXfGawluvnCa", "IsXfGaxwulCa", "IXfGawlqtCva", "IXjfGawlmeCa", "IXfGnaynwlCa", "IXfGcamwelCa"},
		"IXfGawlCa",
		[]bool{true, true, true, true, true, true},
	},

	{
		[]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"},
		"CooP",
		[]bool{false, false, true},
	},

	{
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		"FoBaT",
		[]bool{false, true, false, false, false},
	},

	{
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		"FB",
		[]bool{true, false, true, true, false},
	},

	{
		[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		"FoBa",
		[]bool{true, false, true, false, false},
	},

	// 可以有多个 testcase
}

func Test_camelMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, camelMatch(tc.queries, tc.pattern), "输入:%v", tc)
	}
}

func Benchmark_camelMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			camelMatch(tc.queries, tc.pattern)
		}
	}
}
