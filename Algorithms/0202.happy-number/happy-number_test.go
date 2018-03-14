package problem0202

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans bool
}{

	{
		2,
		false,
	},

	{
		77,
		false,
	},

	{
		19,
		true,
	},

	// 可以有多个 testcase
}

func Test_isHappy(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isHappy(tc.n), "输入:%v", tc)
	}
}

func Benchmark_isHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isHappy(tc.n)
		}
	}
}

func Benchmark_mult(b *testing.B) {
	n := 12345
	for i := 0; i < b.N; i++ {
		_ = (n % 10) * (n % 10)
	}
}
func Benchmark_square(b *testing.B) {
	n := 12345
	for i := 0; i < b.N; i++ {
		_ = square(n % 10)
	}
}

func square(a int) int {
	return a * a
}
