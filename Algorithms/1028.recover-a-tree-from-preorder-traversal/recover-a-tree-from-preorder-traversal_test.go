package problem1028

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans []int
}{

	{
		"3",
		[]int{3},
	},

	{
		"1-2--3--4-5--6--7",
		[]int{1, 2, 5, 3, 4, 6, 7},
	},

	{
		"1-2--3---4-5--6---7",
		[]int{1, 2, 5, 3, kit.NULL, 6, kit.NULL, 4, kit.NULL, 7},
	},

	{
		"1-401--349---90--88",
		[]int{1, 401, kit.NULL, 349, 88, 90},
	},

	// 可以有多个 testcase
}

func Test_recoverFromPreorder(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ans := recoverFromPreorder(tc.S)
		ast.Equal(tc.ans, kit.Tree2ints(ans), "输入:%v", tc)
	}
}

func Benchmark_recoverFromPreorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			recoverFromPreorder(tc.S)
		}
	}
}
