package problem0389

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans byte
}{

	{
		"abcd",
		"abcde",
		'e',
	},

	// 可以有多个 testcase
}

func Test_findTheDifference(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findTheDifference(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_findTheDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findTheDifference(tc.s, tc.t)
		}
	}
}
