package Problem0685

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	edges [][]int
	ans  []int 
}{


	
	// 可以有多个 testcase
}

func Test_findRedundantDirectedConnection(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRedundantDirectedConnection(tc.edges), "输入:%v", tc)
	}
}

func Benchmark_findRedundantDirectedConnection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRedundantDirectedConnection(tc.edges)
		}
	}
}
