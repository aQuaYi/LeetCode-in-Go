package problem0770

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
		"(av * 9) - (ar + 0) - ((bq - cv) + v * (b + bq - bk)) * (a - 12 + 2 - (6 * cc - 8 - bv + ag))",
		[]string{"d", "g", "h", "j", "l", "o", "s", "u", "v", "w", "af", "ag", "ah", "ak", "at", "au", "av", "aw", "az", "bc", "be", "bg", "bj", "bm", "bn", "bq", "br", "bs", "bt", "bu", "bv", "bw", "bx", "by", "bz", "ca", "cd", "ce", "cf", "ch", "ci", "ck", "cq", "cr", "cs", "cu", "cv"},
		[]int{3, 6, 7, 9, 11, 1, 5, 7, 8, 9, 10, 11, 12, 2, 11, 12, 0, 1, 4, 12, 1, 3, 6, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 5, 6, 7, 9, 10, 12, 5, 6, 7, 9, 10},
		[]string{"-8*a*b", "8*a*bk", "48*b*cc", "-48*bk*cc", "10*a", "-1*ar", "64*b", "-64*bk", "-60*cc", "-80"},
	},

	{
		"(e + 8) * (e - 8)",
		[]string{},
		[]int{},
		[]string{"1*e*e", "-64"},
	},

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
