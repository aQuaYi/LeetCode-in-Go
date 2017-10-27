package Problem0335

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x []int
	ans  bool 
}{
	Given x = [2, 1, 1, 2],
	?????
	?   ?
	???????>
		?
	
	Return true (self crossing)
	
	Example 2:
	Given x = [1, 2, 3, 4],
	????????
	?      ?
	?
	?
	?????????????>
	
	Return false (not self crossing)
	
	Example 3:
	Given x = [1, 1, 1, 1],
	?????
	?   ?
	?????>
	
	Return true (self crossing)
	{ },
	{ },

	// 可以有多个 testcase
}

func Test_isSelfCrossing(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isSelfCrossing(tc.x), "输入:%v", tc)
	}
}

func Benchmark_isSelfCrossing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSelfCrossing(tc.x)
		}
	}
}
