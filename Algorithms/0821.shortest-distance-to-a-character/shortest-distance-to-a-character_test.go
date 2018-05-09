package problem0821

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	C   byte
	ans []int
}{

	{
		"abaa",
		'b',
		[]int{1, 0, 1, 2},
	},
	{
		"loveleetcode",
		'e',
		[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
	},

	// 可以有多个 testcase
}

func Test_shortestToChar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestToChar(tc.S, tc.C), "输入:%v", tc)
	}
}

func Benchmark_shortestToChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestToChar(tc.S, tc.C)
		}
	}
}
