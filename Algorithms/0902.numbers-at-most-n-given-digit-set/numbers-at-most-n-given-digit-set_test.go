package problem0902

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	D   []string
	N   int
	ans int
}{

	{
		[]string{"1", "3", "5", "7"},
		100,
		20,
	},

	{
		[]string{"1", "4", "9"},
		1000000000,
		29523,
	},

	// 可以有多个 testcase
}

func Test_atMostNGivenDigitSet(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, atMostNGivenDigitSet(tc.D, tc.N), "输入:%v", tc)
	}
}

func Benchmark_atMostNGivenDigitSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			atMostNGivenDigitSet(tc.D, tc.N)
		}
	}
}
