package problem0959

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid []string
	ans  int
}{

	{
		[]string{
			" /",
			"/ ",
		},
		2,
	},

	{
		[]string{
			"/",
		},
		2,
	},

	{
		[]string{
			" /",
			"  ",
		},
		1,
	},

	{
		[]string{
			"\\/",
			"/\\",
		},
		4,
	},

	{
		[]string{
			"/\\",
			"\\/",
		},
		5,
	},

	{
		[]string{
			"//",
			"/ ",
		},
		3,
	},

	// // 可以有多个 testcase
}

func Test_regionsBySlashes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, regionsBySlashes(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_regionsBySlashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			regionsBySlashes(tc.grid)
		}
	}
}
