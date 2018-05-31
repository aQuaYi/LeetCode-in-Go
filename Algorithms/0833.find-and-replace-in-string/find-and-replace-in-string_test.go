package problem0833

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S       string
	indexes []int
	sources []string
	targets []string
	ans     string
}{

	{
		"abcd",
		[]int{0, 2},
		[]string{"a", "cd"},
		[]string{"eee", "ffff"},
		"eeebffff",
	},

	{
		"abcd",
		[]int{0, 2},
		[]string{"ab", "ec"},
		[]string{"eee", "ffff"},
		"eeecd",
	},

	// 可以有多个 testcase
}

func Test_findReplaceString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findReplaceString(tc.S, tc.indexes, tc.sources, tc.targets), "输入:%v", tc)
	}
}

func Benchmark_findReplaceString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findReplaceString(tc.S, tc.indexes, tc.sources, tc.targets)
		}
	}
}
