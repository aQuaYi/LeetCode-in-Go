package problem0198

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
		[]int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124},
		7102,
	},

	{
		[]int{100, 1, 1, 100, 1, 1, 100, 1, 1, 100},
		400,
	},

	{
		[]int{2, 1, 1, 2},
		4,
	},

	{
		[]int{},
		0,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		20,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		25,
	},

	// 可以有多个 testcase
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, rob(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rob(tc.nums)
		}
	}
}
