package Problem0628

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	temp := append(nums[l-3:], nums[:3]...)

	max := nums[0] * nums[1] * nums[2]

	for i := 0; i < len(temp)-3; i++ {
		t := temp[i] * temp[i+1] * temp[i+2]
		if max < t {
			max = t
		}
	}

	return max
}
