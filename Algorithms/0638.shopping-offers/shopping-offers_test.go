package problem0638

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	price   []int
	special [][]int
	needs   []int
	ans     int
}{

	{
		[]int{2, 5},
		[][]int{{3, 0, 100}, {1, 2, 100}},
		[]int{3, 2},
		16,
	},

	{
		[]int{2, 5},
		[][]int{{3, 0, 5}, {1, 2, 10}},
		[]int{3, 2},
		14,
	},

	{
		[]int{2, 3, 4},
		[][]int{{1, 1, 0, 4}, {2, 2, 1, 9}},
		[]int{1, 2, 1},
		11,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shoppingOffers(tc.price, tc.special, tc.needs), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shoppingOffers(tc.price, tc.special, tc.needs)
		}
	}
}
