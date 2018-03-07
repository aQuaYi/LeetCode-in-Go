package Problem0770

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	evalvars   []string
	evalints   []int
	ans        []string
}{

	{
		"e + 8 - a + 5",
		[]string{"e"},
		[]int{1},
		[]string{"-1*a", "14"},
	},

	{
		"e - 8 + temperature - pressure",
		[]string{"e", "temperature"},
		[]int{1, 12},
		[]string{"-1*pressure", "5"},
	},

	{
		"(e + 8) * (e - 8)",
		[]string{},
		[]int{},
		[]string{"1*e*e", "-64"},
	},

	{
		"7 - 7",
		[]string{},
		[]int{},
		[]string{},
	},

	{
		"a * b * c + b * a * c * 4",
		[]string{},
		[]int{},
		[]string{"5*a*b*c"},
	},

	{
		"((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))",
		[]string{},
		[]int{},
		[]string{"-1*a*a*b*b", "2*a*a*b*c", "-1*a*a*c*c", "1*a*b*b*b", "-1*a*b*b*c", "-1*a*b*c*c", "1*a*c*c*c", "-1*b*b*b*c", "2*b*b*c*c", "-1*b*c*c*c", "2*a*a*b", "-2*a*a*c", "-2*a*b*b", "2*a*c*c", "1*b*b*b", "-1*b*b*c", "1*b*c*c", "-1*c*c*c", "-1*a*a", "1*a*b", "1*a*c", "-1*b*c"},
	},

	// 可以有多个 testcase
}

func Test_basicCalculatorIV(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, basicCalculatorIV(tc.expression, tc.evalvars, tc.evalints), "输入:%v", tc)
	}
}

func Benchmark_basicCalculatorIV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			basicCalculatorIV(tc.expression, tc.evalvars, tc.evalints)
		}
	}
}
