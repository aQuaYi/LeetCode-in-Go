package problem0164

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 自己实现了一个简易的快排
	var quickSort func(i, j int)
	quickSort = func(i, j int) {
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
			quickSort(ci, c)
		}
		if c+1 < cj {
			quickSort(c+1, cj)
		}
	}

	quickSort(0, len(nums)-1)

	ret := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if ret < tmp {
			ret = tmp
		}
	}

	return ret
}
