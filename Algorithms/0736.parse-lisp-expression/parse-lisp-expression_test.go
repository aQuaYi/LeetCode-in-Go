package problem0736

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	ans        int
}{

	{
		"(let x0 4 x1 -2 x2 3 x3 -5 x4 -3 x5 -1 x6 3 x7 -2 x8 4 x9 -5 (mult x2 (mult (let x0 -3 x4 -2 x8 4 (mult (let x0 -2 x6 4 (add x5 x2)) x3)) (mult (mult -7 (mult -9 (let x0 -2 x7 3 (add -10 x0)))) x6))))",
		68040,
	},

	{
		"(mult 3 (add 2 3))",
		15,
	},

	{
		"(add 1 2)",
		3,
	},

	{
		"(let x 2 (mult x 5))",
		10,
	},

	{
		"(let x 2 (mult x (let x 3 y 4 (add x y))))",
		14,
	},

	{
		"(let x 3 x 2 x)",
		2,
	},

	{
		"(let x 1 y 2 x (add x y) (add x y))",
		5,
	},

	{
		"(let x 2 (add (let x 3 (let x 4 x)) x))",
		6,
	},

	{
		"(let a1 3 b2 (add a1 1) b2)",
		4,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, evaluate(tc.expression), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			evaluate(tc.expression)
		}
	}
}
