package problem0279

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{
		123456,
		3,
	},

	{
		16,
		1,
	},

	{
		15,
		4,
	},

	{
		14,
		3,
	},

	{
		13,
		2,
	},

	{
		12,
		3,
	},

	// 可以有多个 testcase
}

func Test_numSquares(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numSquares(tc.n), "输入:%v", tc)
	}
}

func Benchmark_numSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSquares(tc.n)
		}
	}
}

var max = 10000000

func Benchmark_math_sqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10; j <= max; j++ {
			_ = int(math.Sqrt(float64(j)))
		}
	}
}

func Benchmark_intSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10; j <= max; j *= 10 {
			_ = intSqrt(j)
		}
	}
}

// 返回 x 的平方根的整数部分
// 这个函数比 int(math.Sqrt(float64(x))) 快的多
// 详见 benchmark 的结果
func intSqrt(x int) int {
	res := x

	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
