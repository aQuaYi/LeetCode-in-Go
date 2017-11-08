package Problem0435

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals []Interval
	ans  int 
}{


	
	// 可以有多个 testcase
}

func Test_eraseOverlapIntervals(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, eraseOverlapIntervals(tc.intervals), "输入:%v", tc)
	}
}

func Benchmark_eraseOverlapIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			eraseOverlapIntervals(tc.intervals)
		}
	}
}
