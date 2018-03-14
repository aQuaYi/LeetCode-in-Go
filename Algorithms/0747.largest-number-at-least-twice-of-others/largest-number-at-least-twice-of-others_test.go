package problem0747

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1},
		0,
	},

	{
		[]int{0, 3, 1, 1},
		1,
	},

	{
		[]int{0, 0, 0, 1},
		3,
	},

	{
		[]int{3, 6, 1, 0},
		1,
	},

	{
		[]int{1, 2, 3, 4},
		-1,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, dominantIndex(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			dominantIndex(tc.nums)
		}
	}
}
