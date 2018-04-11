package problem0791

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	T   string
	ans string
}{

	{
		"cba",
		"abcd",
		"cbad",
	},

	{
		"cba",
		"abcdbca",
		"ccbbaad",
	},

	// 可以有多个 testcase
}

func Test_customSortString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, customSortString(tc.S, tc.T), "输入:%v", tc)
	}
}

func Benchmark_customSortString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			customSortString(tc.S, tc.T)
		}
	}
}
