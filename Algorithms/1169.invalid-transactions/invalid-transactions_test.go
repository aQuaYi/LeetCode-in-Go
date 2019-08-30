package problem1169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	transactions []string
	ans          []string
}{

	{
		[]string{"alice,120,800,mtv", "alice,20,800,mtv", "alice,50,100,beijing"},
		[]string{"alice,20,800,mtv", "alice,50,100,beijing"},
	},

	{
		[]string{"alice,20,800,mtv", "alice,50,100,beijing"},
		[]string{"alice,20,800,mtv", "alice,50,100,beijing"},
	},

	{
		[]string{"alice,20,800,mtv", "alice,50,1200,mtv"},
		[]string{"alice,50,1200,mtv"},
	},

	{
		[]string{"alice,20,800,mtv", "bob,50,1200,mtv"},
		[]string{"bob,50,1200,mtv"},
	},

	// 可以有多个 testcase
}

func Test_invalidTransactions(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, invalidTransactions(tc.transactions), "输入:%v", tc)
	}
}

func Benchmark_invalidTransactions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			invalidTransactions(tc.transactions)
		}
	}
}
