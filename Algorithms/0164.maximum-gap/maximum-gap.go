package Problem0164

func maximumGap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var sort func(i, j int)
	sort = func(i, j int) {
		ci, cj := i, j
		c := i
		for i < j {
			for nums[j] >= nums[c] && i < j {
				j--
			}
			for nums[i] <= nums[c] && i < j {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[c] = nums[c], nums[i]
		c = i
		if ci < c {
			sort(ci, c)
		}
		if c+1 < cj {
			sort(c+1, cj)
		}
	}
	sort(0, len(nums)-1)

	ret := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if ret < tmp {
			ret = tmp
		}
	}

	return ret
}
