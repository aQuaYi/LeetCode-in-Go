package problem0461

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   int
	y   int
	ans int
}{

	{
		3,
		1,
		1,
	},

	{
		1,
		4,
		2,
	},

	// 可以有多个 testcase
}

func Test_hammingDistance(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, hammingDistance(tc.x, tc.y), "输入:%v", tc)
	}
}

func Benchmark_hammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hammingDistance(tc.x, tc.y)
		}
	}
}
