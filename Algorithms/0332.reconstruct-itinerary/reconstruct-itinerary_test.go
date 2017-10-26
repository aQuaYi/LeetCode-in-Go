package Problem0332

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tickets [][]string
	ans  []string 
}{
	{[][]string{[]string{ "MUC", "LHR" }, []string{"JFK", "MUC"}, []string{"SFO", "SJC"}, []string{"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
	
	{[][]string{[]string{"JFK","SFO"},[]string{"JFK","ATL"},[]string{"SFO","ATL"},[]string{"ATL","JFK"},[]string{"ATL","SFO"}}, []string{"JFK","ATL","JFK","SFO","ATL","SFO"}},

	// 可以有多个 testcase
}

func Test_findItinerary(t *testing.T) {
	ast := assert.New(t)
	
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findItinerary(tc.tickets), "输入:%v", tc)
	}
}

func Benchmark_findItinerary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findItinerary(tc.tickets)
		}
	}
}
