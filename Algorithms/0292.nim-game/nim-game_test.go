package problem0292

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans bool
}{

	{4, false},
	{5, true},
	{41, true},

	// 可以有多个 testcase
}

func Test_canWinNim(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canWinNim(tc.n), "输入:%v", tc)
	}

}

func Benchmark_canWinNim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canWinNim(tc.n)
		}
	}
}
