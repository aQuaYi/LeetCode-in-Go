package problem0396

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{[]int{4, 3, 2, 6}, 26},

	// 可以有多个 testcase
}

func Test_maxRotateFunction(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxRotateFunction(tc.A), "输入:%v", tc)
	}
}

func Benchmark_maxRotateFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxRotateFunction(tc.A)
		}
	}
}
