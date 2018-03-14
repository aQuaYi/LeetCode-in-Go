package problem0324

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)

	mid := n / 2

	for i := 0; i < n; i++ {
		if i < mid {
			// 把较大的数，按照降序插入 nums 的奇数位
			nums[2*i+1] = temp[n-1-i]
		} else {
			// 把较小的数，按照降序插入 nums 的偶数位
			nums[2*(i-mid)] = temp[n-1-i]
		}
	}
}
