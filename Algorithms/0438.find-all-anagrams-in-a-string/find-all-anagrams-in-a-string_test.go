package problem0438

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	p   string
	ans []int
}{

	{
		"abc",
		"cbaebabacd",
		[]int{},
	},

	{
		"cbaebabacd",
		"abc",
		[]int{0, 6},
	},

	{
		"abab",
		"ab",
		[]int{0, 1, 2},
	},

	// 可以有多个 testcase
}

func Test_findAnagrams(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findAnagrams(tc.s, tc.p), "输入:%v", tc)
	}
}

func Benchmark_findAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findAnagrams(tc.s, tc.p)
		}
	}
}
