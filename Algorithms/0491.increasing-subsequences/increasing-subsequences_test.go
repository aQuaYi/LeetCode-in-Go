package Problem0491

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  [][]int
}{

	{
		[]int{4, 6, 7, 7},
		[][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}},
	},

	// 可以有多个 testcase
}

func Test_findSubsequences(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findSubsequences(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_findSubsequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findSubsequences(tc.nums)
		}
	}
}
