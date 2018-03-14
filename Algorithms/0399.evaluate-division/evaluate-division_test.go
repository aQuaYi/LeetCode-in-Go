package problem0399

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	equations [][]string
	values    []float64
	queries   [][]string
	ans       []float64
}{

	{
		[][]string{[]string{"x1", "x2"}, []string{"x2", "x3"}, []string{"x1", "x4"}, []string{"x2", "x5"}},
		[]float64{3.0, 0.5, 3.4, 5.6},
		[][]string{[]string{"x2", "x4"}, []string{"x1", "x5"}, []string{"x1", "x3"}, []string{"x5", "x5"}, []string{"x5", "x1"}, []string{"x3", "x4"}, []string{"x4", "x3"}, []string{"x6", "x6"}, []string{"x0", "x0"}},
		[]float64{1.13333, 16.80000, 1.50000, 1.00000, 0.05952, 2.26667, 0.44118, -1.00000, -1.00000},
	},

	{
		[][]string{[]string{"a", "b"}, []string{"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}},
		[]float64{6.0, 0.5, -1.0, 1.0, -1.0},
	},

	{
		[][]string{[]string{"a", "b"}, []string{"e", "f"}, []string{"b", "e"}},
		[]float64{3.4, 1.4, 2.3},
		[][]string{[]string{"b", "a"}, []string{"a", "f"}, []string{"f", "f"}, []string{"e", "e"}, []string{"c", "c"}, []string{"a", "c"}, []string{"f", "e"}},
		[]float64{0.29412, 10.94800, 1.00000, 1.00000, -1.00000, -1.00000, 0.71429},
	},

	{
		[][]string{[]string{"a", "b"}, []string{"c", "d"}},
		[]float64{1.0, 1.0},
		[][]string{[]string{"a", "c"}, []string{"b", "d"}, []string{"b", "a"}, []string{"d", "c"}},
		[]float64{-1.00000, -1.00000, 1.00000, 1.00000},
	},

	// 可以有多个 testcase
}

func Test_calcEquation(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := calcEquation(tc.equations, tc.values, tc.queries)
		for i := range ans {
			ast.InDelta(tc.ans[i], ans[i], 0.0001)
		}
	}
}

func Benchmark_calcEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calcEquation(tc.equations, tc.values, tc.queries)
		}
	}
}
