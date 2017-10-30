package Problem0391

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rectangles [][]int
	ans  bool 
}{


	
	// 可以有多个 testcase
}

func Test_isRectangleCover(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isRectangleCover(tc.rectangles), "输入:%v", tc)
	}
}

func Benchmark_isRectangleCover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isRectangleCover(tc.rectangles)
		}
	}
}
