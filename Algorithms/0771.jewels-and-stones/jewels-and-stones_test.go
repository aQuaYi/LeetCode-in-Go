package problem0771

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	J   string
	S   string
	ans int
}{

	{
		"aA",
		"aAAbbbb",
		3,
	},

	{
		"z",
		"ZZ",
		0,
	},

	// 可以有多个 testcase
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numJewelsInStones(tc.J, tc.S), "输入:%v", tc)
	}
}

func Benchmark_numJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numJewelsInStones(tc.J, tc.S)
		}
	}
}
