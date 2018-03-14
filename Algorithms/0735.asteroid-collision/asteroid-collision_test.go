package problem0735

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	asteroids []int
	ans       []int
}{

	{
		[]int{5, 10, -5},
		[]int{5, 10},
	},

	{
		[]int{8, -8},
		[]int{},
	},

	{
		[]int{10, 2, -5},
		[]int{10},
	},

	{
		[]int{-2, -1, 1, 2},
		[]int{-2, -1, 1, 2},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, asteroidCollision(tc.asteroids), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			asteroidCollision(tc.asteroids)
		}
	}
}
