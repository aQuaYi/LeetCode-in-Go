package problem0365

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   int
	y   int
	z   int
	ans bool
}{

	{11, 3, 13, true},
	{0, 0, 0, true},
	{3, 5, 4, true},
	{4, 6, 8, true},
	{3, 5, 1, true},
	{2, 6, 5, false},
	{2, 6, 15, false},
	{0, 2, 1, false},

	// 可以有多个 testcase
}

func Test_canMeasureWater(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canMeasureWater(tc.x, tc.y, tc.z), "输入:%v", tc)
	}
}

func Benchmark_canMeasureWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canMeasureWater(tc.x, tc.y, tc.z)
		}
	}
}
