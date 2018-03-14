package problem0274

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	citations []int
	ans       int
}{

	{
		[]int{3, 0, 6, 1, 5},
		3,
	},

	// 可以有多个 testcase
}

func Test_hIndex(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, hIndex(tc.citations), "输入:%v", tc)
	}
}

func Benchmark_hIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hIndex(tc.citations)
		}
	}
}
