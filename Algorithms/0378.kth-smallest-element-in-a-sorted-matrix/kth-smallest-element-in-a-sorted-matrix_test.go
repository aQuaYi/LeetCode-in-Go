package problem0378

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mat = [][]int{
	[]int{1, 5, 9},
	[]int{10, 11, 13},
	[]int{12, 13, 15},
}

// tcs is testcase slice
var tcs = []struct {
	k   int
	ans int
}{

	{3, 9},

	{8, 13},

	// 可以有多个 testcase
}

func Test_kthSmallest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, kthSmallest(mat, tc.k), "输入:%v", tc)
	}
}

func Benchmark_kthSmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			kthSmallest(mat, tc.k)
		}
	}
}
