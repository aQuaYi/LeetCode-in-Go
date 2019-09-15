package problem1185

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	day   int
	month int
	year  int
	ans   string
}{

	{
		31,
		8,
		2019,
		"Saturday",
	},

	{
		18,
		7,
		1999,
		"Sunday",
	},

	{
		15,
		8,
		1993,
		"Sunday",
	},

	// 可以有多个 testcase
}

func Test_dayOfTheWeek(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, dayOfTheWeek(tc.day, tc.month, tc.year), "输入:%v", tc)
	}
}

func Benchmark_dayOfTheWeek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			dayOfTheWeek(tc.day, tc.month, tc.year)
		}
	}
}
