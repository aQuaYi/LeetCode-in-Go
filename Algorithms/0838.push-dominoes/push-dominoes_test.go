package problem0838

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dominoes string
	ans      string
}{

	{
		".L.R...LR..L..",
		"LL.RR.LLRRLL..",
	},

	{
		"RR.L",
		"RR.L",
	},

	// 可以有多个 testcase
}

func Test_pushDominoes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, pushDominoes(tc.dominoes), "输入:%v", tc)
	}
}

func Benchmark_pushDominoes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pushDominoes(tc.dominoes)
		}
	}
}
