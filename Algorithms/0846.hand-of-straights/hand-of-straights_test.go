package problem0846

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	hand []int
	W    int
	ans  bool
}{

	{
		[]int{1, 2, 3},
		1,
		true,
	},

	{
		[]int{1, 2, 3, 7, 2, 3, 4, 7, 8},
		3,
		false,
	},

	{
		[]int{1, 2, 3, 6, 2, 3, 4, 7, 8},
		3,
		true,
	},

	{
		[]int{1, 2, 3, 4, 5},
		4,
		false,
	},

	// 可以有多个 testcase
}

func Test_isNStraightHand(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isNStraightHand(tc.hand, tc.W), "输入:%v", tc)
	}
}

func Benchmark_isNStraightHand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isNStraightHand(tc.hand, tc.W)
		}
	}
}
