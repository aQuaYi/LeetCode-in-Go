package Problem0324

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	var i, j int
	// 把 大数 按照降序，排列在奇数位
	i = 0
	j = 1
	for j < n {
		nums[j] = temp[i]
		i++
		j += 2
	}
	// 把 小数 按照降序，排列在偶数位
	i, j = n-1, n-1
	if j%2 != 0 {
		j--
	}
	for 0 <= j {
		nums[j] = temp[i]
		j -= 2
		i--
	}
}
