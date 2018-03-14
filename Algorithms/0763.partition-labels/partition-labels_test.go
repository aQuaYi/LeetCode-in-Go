package problem0763

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []int
}{

	{
		"abcdefg",
		[]int{1, 1, 1, 1, 1, 1, 1},
	},

	{
		"ababcbacadefegdehijhklij",
		[]int{9, 7, 8},
	},

	// 可以有多个 testcase
}

func Test_partitionLabels(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, partitionLabels(tc.S), "输入:%v", tc)
	}
}

func Benchmark_partitionLabels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			partitionLabels(tc.S)
		}
	}
}
