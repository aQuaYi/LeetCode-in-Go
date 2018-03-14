package problem0393

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	data []int
	ans  bool
}{

	{
		[]int{241, 130, 130, 130, 1},
		true,
	},

	{
		[]int{197, 130, 1},
		true,
	},

	{
		[]int{129},
		false,
	},

	{
		[]int{235, 140, 4},
		false,
	},

	// 可以有多个 testcase
}

func Test_validUtf8(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, validUtf8(tc.data), "输入:%v", tc)
	}
}

func Benchmark_validUtf8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validUtf8(tc.data)
		}
	}
}
