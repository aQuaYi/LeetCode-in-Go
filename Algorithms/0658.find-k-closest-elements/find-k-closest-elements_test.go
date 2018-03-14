package problem0658

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr []int
	k   int
	x   int
	ans []int
}{

	{
		[]int{1, 2, 3, 4, 5},
		4,
		3,
		[]int{1, 2, 3, 4},
	},

	{
		[]int{1, 2, 3, 4, 5},
		4,
		-1,
		[]int{1, 2, 3, 4},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findClosestElements(tc.arr, tc.k, tc.x), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findClosestElements(tc.arr, tc.k, tc.x)
		}
	}
}
