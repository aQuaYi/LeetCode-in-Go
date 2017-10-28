package Problem0385

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n NestedInteger
	ans  IsInteger() bool 
}{


	
	// 可以有多个 testcase
}

func Test_(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, (tc.n), "输入:%v", tc)
	}
}

func Benchmark_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			(tc.n)
		}
	}
}
