package problem0970

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x     int
	y     int
	bound int
	ans   []int
}{

	{
		1,
		1,
		0,
		[]int{},
	},

	{
		2,
		1,
		10,
		[]int{2, 3, 5, 9},
	},

	{
		2,
		3,
		10,
		[]int{2, 3, 4, 5, 7, 9, 10},
	},

	{
		3,
		5,
		15,
		[]int{2, 4, 6, 8, 10, 14},
	},

	// 可以有多个 testcase
}

func Test_powerfulIntegers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		sort.Ints(tc.ans)
		ans := powerfulIntegers(tc.x, tc.y, tc.bound)
		sort.Ints(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_powerfulIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			powerfulIntegers(tc.x, tc.y, tc.bound)
		}
	}
}

func Benchmark_removeRepeated(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}
	for i := 1; i < b.N; i++ {
		removeRepeated(nums)
	}
}

func Benchmark_removeRepeated2(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}
	for i := 1; i < b.N; i++ {
		removeRepeated2(nums)
	}
}

func removeRepeated2(nums []int) []int {
	sort.Ints(nums)

	size := len(nums)
	if size == 0 {
		return nums
	}

	i := 0
	for j := 1; j < size; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return nums[:i+1]
}
