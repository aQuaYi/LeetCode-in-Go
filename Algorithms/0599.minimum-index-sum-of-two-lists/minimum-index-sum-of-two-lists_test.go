package problem0599

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	list1, list2, ans []string
}{

	{
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		[]string{"Shogun"},
	},

	{
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
		[]string{"Shogun"},
	},

	{
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
	},

	// 可以有多个 testcase
}

func Test_fcName(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRestaurant(tc.list1, tc.list2), "输入:%v", tc)
	}
}

func Benchmark_fcName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRestaurant(tc.list1, tc.list2)
		}
	}
}
