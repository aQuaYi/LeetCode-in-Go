package problem0863

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root   []int
	target []int
	K      int
	ans    []int
}{

	{
		[]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
		5,
		2,
		[]int{7, 4, 1},
	},

	// 可以有多个 testcase
}

func Test_distanceK(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, distanceK(tc.root, tc.target, tc.K), "输入:%v", tc)
	}
}

func Benchmark_distanceK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			distanceK(tc.root, tc.target, tc.K)
		}
	}
}
