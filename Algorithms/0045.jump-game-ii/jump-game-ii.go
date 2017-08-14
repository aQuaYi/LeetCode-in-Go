package Problem0045

func jump(nums []int) int {
	res := len(nums)

	jumpToNext(nums, &res, 0, 0)

	return res
}

func jumpToNext(nums []int, res *int, i, temp int) {
	if temp > *res {
		return
	}

	if i >= len(nums)-1 {
		if *res > temp {
			*res = temp
		}
		return
	}

	for j := nums[i]; j > 0; j-- {
		jumpToNext(nums, res, i+j, temp+1)
	}
}
