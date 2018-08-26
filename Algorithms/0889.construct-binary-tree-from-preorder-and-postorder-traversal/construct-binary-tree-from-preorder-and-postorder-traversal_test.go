package problem0889

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre []int
 post []int
	ans *TreeNode
}{


	
	// 可以有多个 testcase
}

func Test_constructFromPrePost(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, constructFromPrePost(tc.pre, tc.post), "输入:%v", tc)
	}
}

func Benchmark_constructFromPrePost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			constructFromPrePost(tc.pre, tc.post)
		}
	}
}
