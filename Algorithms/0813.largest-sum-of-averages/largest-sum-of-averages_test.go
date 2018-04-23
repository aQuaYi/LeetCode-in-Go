package problem0813

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans float64
}{

	{
		[]int{9, 1, 2, 3, 9},
		3,
		20,
	},

	{
		[]int{4663, 3020, 7789, 1627, 9668, 1356, 4207, 1133, 8765, 4649, 205, 6455, 8864, 3554, 3916, 5925, 3995, 4540, 3487, 5444, 8259, 8802, 6777, 7306, 989, 4958, 2921, 8155, 4922, 2469, 6923, 776, 9777, 1796, 708, 786, 3158, 7369, 8715, 2136, 2510, 3739, 6411, 7996, 6211, 8282, 4805, 236, 1489, 7698},
		27,
		167436.08333,
	},

	// 可以有多个 testcase
}

func Test_largestSumOfAverages(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.InDelta(tc.ans, largestSumOfAverages(tc.A, tc.K), 0.00001, "输入:%v", tc)
	}
}

func Benchmark_largestSumOfAverages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestSumOfAverages(tc.A, tc.K)
		}
	}
}
