package Problem0324

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)

	mid := n / 2

	for i := 0; i < n; i++ {
		if i < mid {
			nums[2*i+1] = temp[n-1-i]
		} else {
			nums[2*(i-mid)] = temp[n-1-i]
		}
	}
}
