package problem0377

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    int
}{

	{
		[]int{1, 2, 3},
		4,
		7,
	},

	// 可以有多个 testcase
}

func Test_combinationSum4(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, combinationSum4(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_combinationSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			combinationSum4(tc.nums, tc.target)
		}
	}
}
