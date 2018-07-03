package problem0848

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S      string
	shifts []int
	ans    string
}{

	{
		"abc",
		[]int{3, 5, 9},
		"rpl",
	},

	// 可以有多个 testcase
}

func Test_shiftingLetters(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shiftingLetters(tc.S, tc.shifts), "输入:%v", tc)
	}
}

func Benchmark_shiftingLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shiftingLetters(tc.S, tc.shifts)
		}
	}
}
