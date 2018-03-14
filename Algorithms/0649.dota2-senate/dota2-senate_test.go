package problem0649

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	senate string
	ans    string
}{

	{
		"RRDDD",
		"Radiant",
	},

	{
		"DDRRR",
		"Dire",
	},

	{
		"RD",
		"Radiant",
	},

	{
		"RDD",
		"Dire",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, predictPartyVictory(tc.senate), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			predictPartyVictory(tc.senate)
		}
	}
}
