package problem1189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text string
	ans  int
}{

	{
		"nlaeolko",
		0,
	},

	{
		"nlaebolko",
		1,
	},

	{
		"loonbalxballpoon",
		2,
	},

	{
		"leetcode",
		0,
	},

	// 可以有多个 testcase
}

func Test_maxNumberOfBalloons(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxNumberOfBalloons(tc.text), "输入:%v", tc)
	}
}

func Benchmark_maxNumberOfBalloons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxNumberOfBalloons(tc.text)
		}
	}
}
