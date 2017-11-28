package Problem0501

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []int 
}{


	
	// 可以有多个 testcase
}

func Test_findMode(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMode(tc.root), "输入:%v", tc)
	}
}

func Benchmark_findMode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMode(tc.root)
		}
	}
}
