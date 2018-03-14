package problem0717

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	bits []int
	ans  bool
}{

	{
		[]int{1, 1, 0, 0, 0, 0},
		true,
	},

	{
		[]int{1, 1, 1, 1, 0},
		true,
	},

	{
		[]int{1, 0, 0},
		true,
	},

	{
		[]int{1, 1, 1, 0},
		false,
	},

	{
		[]int{1, 1, 0},
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isOneBitCharacter(tc.bits), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isOneBitCharacter(tc.bits)
		}
	}
}
