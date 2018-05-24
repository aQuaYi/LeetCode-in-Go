package problem0830

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans [][]int
}{

	{
		"abbxxxxzzy",
		[][]int{{3, 6}},
	},

	{
		"abc",
		[][]int{},
	},

	{
		"abcdddeeeeaabbbcd",
		[][]int{{3, 5}, {6, 9}, {12, 14}},
	},

	// 可以有多个 testcase
}

func Test_largeGroupPositions(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, largeGroupPositions(tc.S), "输入:%v", tc)
	}
}

func Benchmark_largeGroupPositions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largeGroupPositions(tc.S)
		}
	}
}
