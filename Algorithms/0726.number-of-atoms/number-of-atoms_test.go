package problem0726

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	formula string
	ans     string
}{

	{
		"H2O",
		"H2O",
	},

	{
		"Mg(OH)2",
		"H2MgO2",
	},

	{
		"K4(ON(SO3)2)2K4(ON(SO3)2)2",
		"K8N4O28S8",
	},

	{
		"K4(ON(SO3)2)2",
		"K4N2O14S4",
	},

	// 可以有多个 testcase
}

func Test_countOfAtoms(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countOfAtoms(tc.formula), "输入:%v", tc)
	}
}

func Benchmark_countOfAtoms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countOfAtoms(tc.formula)
		}
	}
}
