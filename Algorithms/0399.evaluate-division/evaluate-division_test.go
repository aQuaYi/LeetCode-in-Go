package Problem0399

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
		ast.Equal(tc.ans, calcEquation(tc.equations, tc.values, tc.queries), "输入:%v", tc)
	}
}

func Benchmark_calcEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calcEquation(tc.equations, tc.values, tc.queries)
		}
	}
}
