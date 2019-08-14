package problem1154

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	date string
	ans  int
}{

	{
		"2019-01-09",
		9,
	},

	{
		"2019-02-10",
		41,
	},

	{
		"2003-03-01",
		60,
	},

	{
		"2004-03-01",
		61,
	},

	// 可以有多个 testcase
}

func Test_dayOfYear(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, dayOfYear(tc.date), "输入:%v", tc)
	}
}

func Benchmark_dayOfYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			dayOfYear(tc.date)
		}
	}
}
