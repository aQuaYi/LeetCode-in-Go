package problem0952

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{4, 6, 15, 35},
		4,
	},

	{
		[]int{20, 50, 9, 63},
		2,
	},

	{
		[]int{2, 3, 6, 7, 4, 12, 21, 39},
		8,
	},

	{
		[]int{2, 7, 522, 526, 535, 26, 944, 35, 519, 45, 48, 567, 266, 68, 74, 591, 81, 86, 602, 93, 610, 621, 111, 114, 629, 641, 131, 651, 142, 659, 669, 161, 674, 163, 180, 187, 190, 194, 195, 206, 207, 218, 737, 229, 240, 757, 770, 260, 778, 270, 272, 785, 274, 290, 291, 292, 296, 810, 816, 314, 829, 833, 841, 349, 880, 369, 147, 897, 387, 390, 905, 405, 406, 407, 414, 416, 417, 425, 938, 429, 432, 926, 959, 960, 449, 963, 966, 929, 457, 463, 981, 985, 79, 487, 1000, 494, 508},
		84,
	},

	// 可以有多个 testcase
}

func Test_largestComponentSize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, largestComponentSize(tc.A), "输入:%v", tc)
	}
}

func Benchmark_largestComponentSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestComponentSize(tc.A)
		}
	}
}
