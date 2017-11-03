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
