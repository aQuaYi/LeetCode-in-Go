package problem0810

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	name string
	nums []int
	ans  bool
}{

	{
		"1",
		[]int{7, 8, 6, 10, 14, 15, 16, 6, 4, 0, 4, 9, 3, 3, 8, 10, 5, 5, 3, 10},
		true,
	},

	{
		"2",
		[]int{0, 4, 4, 0, 4, 4, 2, 2, 2, 3, 0, 3, 1, 3, 2, 4, 2, 0, 1, 3},
		true,
	},

	{
		"3",
		[]int{1, 1, 2},
		false,
	},

	{
		"4",
		[]int{1, 3, 2},
		true,
	},

	// 可以有多个 testcase
}

func Test_xorGame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, xorGame(tc.nums), "%s, 输入:%v", tc.name, tc)
	}
}

func Benchmark_xorGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			xorGame(tc.nums)
		}
	}
}
