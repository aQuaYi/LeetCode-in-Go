package problem0332

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	tickets [][]string
	ans     []string
}{

	{
		[][]string{
			[]string{"JFK", "SFO"},
			[]string{"JFK", "ATL"},
			[]string{"SFO", "ATL"},
			[]string{"ATL", "JFK"},
			[]string{"ATL", "SFO"},
		},
		[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	},

	{
		[][]string{
			[]string{"MUC", "LHR"},
			[]string{"JFK", "MUC"},
			[]string{"SFO", "SJC"},
			[]string{"LHR", "SFO"},
		},
		[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},

	{
		[][]string{
			[]string{"AXA", "EZE"},
			[]string{"EZE", "AUA"},
			[]string{"ADL", "JFK"},
			[]string{"ADL", "TIA"},
			[]string{"AUA", "AXA"},
			[]string{"EZE", "TIA"},
			[]string{"EZE", "TIA"},
			[]string{"AXA", "EZE"},
			[]string{"EZE", "ADL"},
			[]string{"ANU", "EZE"},
			[]string{"TIA", "EZE"},
			[]string{"JFK", "ADL"},
			[]string{"AUA", "JFK"},
			[]string{"JFK", "EZE"},
			[]string{"EZE", "ANU"},
			[]string{"ADL", "AUA"},
			[]string{"ANU", "AXA"},
			[]string{"AXA", "ADL"},
			[]string{"AUA", "JFK"},
			[]string{"EZE", "ADL"},
			[]string{"ANU", "TIA"},
			[]string{"AUA", "JFK"},
			[]string{"TIA", "JFK"},
			[]string{"EZE", "AUA"},
			[]string{"AXA", "EZE"},
			[]string{"AUA", "ANU"},
			[]string{"ADL", "AXA"},
			[]string{"EZE", "ADL"},
			[]string{"AUA", "ANU"},
			[]string{"AXA", "EZE"},
			[]string{"TIA", "AUA"},
			[]string{"AXA", "EZE"},
			[]string{"AUA", "SYD"},
			[]string{"ADL", "JFK"},
			[]string{"EZE", "AUA"},
			[]string{"ADL", "ANU"},
			[]string{"AUA", "TIA"},
			[]string{"ADL", "EZE"},
			[]string{"TIA", "JFK"},
			[]string{"AXA", "ANU"},
			[]string{"JFK", "AXA"},
			[]string{"JFK", "ADL"},
			[]string{"ADL", "EZE"},
			[]string{"AXA", "TIA"},
			[]string{"JFK", "AUA"},
			[]string{"ADL", "EZE"},
			[]string{"JFK", "ADL"},
			[]string{"ADL", "AXA"},
			[]string{"TIA", "AUA"},
			[]string{"AXA", "JFK"},
			[]string{"ADL", "AUA"},
			[]string{"TIA", "JFK"},
			[]string{"JFK", "ADL"},
			[]string{"JFK", "ADL"},
			[]string{"ANU", "AXA"},
			[]string{"TIA", "AXA"},
			[]string{"EZE", "JFK"},
			[]string{"EZE", "AXA"},
			[]string{"ADL", "TIA"},
			[]string{"JFK", "AUA"},
			[]string{"TIA", "EZE"},
			[]string{"EZE", "ADL"},
			[]string{"JFK", "ANU"},
			[]string{"TIA", "AUA"},
			[]string{"EZE", "ADL"},
			[]string{"ADL", "JFK"},
			[]string{"ANU", "AXA"},
			[]string{"AUA", "AXA"},
			[]string{"ANU", "EZE"},
			[]string{"ADL", "AXA"},
			[]string{"ANU", "AXA"},
			[]string{"TIA", "ADL"},
			[]string{"JFK", "ADL"},
			[]string{"JFK", "TIA"},
			[]string{"AUA", "ADL"},
			[]string{"AUA", "TIA"},
			[]string{"TIA", "JFK"},
			[]string{"EZE", "JFK"},
			[]string{"AUA", "ADL"},
			[]string{"ADL", "AUA"},
			[]string{"EZE", "ANU"},
			[]string{"ADL", "ANU"},
			[]string{"AUA", "AXA"},
			[]string{"AXA", "TIA"},
			[]string{"AXA", "TIA"},
			[]string{"ADL", "AXA"},
			[]string{"EZE", "AXA"},
			[]string{"AXA", "JFK"},
			[]string{"JFK", "AUA"},
			[]string{"ANU", "ADL"},
			[]string{"AXA", "TIA"},
			[]string{"ANU", "AUA"},
			[]string{"JFK", "EZE"},
			[]string{"AXA", "ADL"},
			[]string{"TIA", "EZE"},
			[]string{"JFK", "AXA"},
			[]string{"AXA", "ADL"},
			[]string{"EZE", "AUA"},
			[]string{"AXA", "ANU"},
			[]string{"ADL", "EZE"},
			[]string{"AUA", "EZE"},
		},
		[]string{"JFK", "ADL", "ANU", "ADL", "ANU", "AUA", "ADL", "AUA", "ADL", "AUA", "ANU", "AXA", "ADL", "AUA", "ANU", "AXA", "ADL", "AXA", "ADL", "AXA", "ANU", "AXA", "ANU", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "AXA", "EZE", "ADL", "EZE", "ADL", "EZE", "ADL", "EZE", "ANU", "EZE", "ANU", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "EZE", "AUA", "AXA", "JFK", "ADL", "EZE", "AUA", "EZE", "AXA", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "JFK", "ADL", "TIA", "ADL", "TIA", "AUA", "JFK", "ANU", "TIA", "AUA", "JFK", "AUA", "JFK", "AUA", "TIA", "AUA", "TIA", "AXA", "TIA", "EZE", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "EZE", "JFK", "AXA", "TIA", "JFK", "EZE", "TIA", "JFK", "EZE", "TIA", "JFK", "TIA", "JFK", "AUA", "SYD"},
	},

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
