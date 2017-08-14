package Problem0045

func jump(nums []int) int {
	ress := []int{}

	jumpToNext(nums, &ress, 0, 0)

	min := len(nums) + 1
	for i := range ress {
		if min > ress[i] {
			min = ress[i]
		}
	}

	return min
}

func jumpToNext(nums []int, ress *[]int, i, temp int) {
	if nums[i]+i >= len(nums)-1 {
		*ress = append(*ress, temp+1)
		return
	}

	for j := 1; j <= nums[i]; j++ {
		jumpToNext(nums, ress, i+j, temp+1)
	}

}
