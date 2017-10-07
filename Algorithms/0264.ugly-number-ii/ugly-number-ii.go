package Problem0264

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	nums := make([]int, 6, n)
	for i := 0; i < 6; i++ {
		nums[i] = i + 1
	}

	var search func(int, int, int)
	search = func(left, right, max int) {
		if left > right {
			return
		}
		min := nums[len(nums)-1]
		temp := 0

		for left <= right {
			temp = nums[left] * nums[right]

			if max <= temp {
				right--
			} else if temp <= min {
				left++
			} else {
				search(left, right-1, temp)
				if len(nums) < n {
					nums = append(nums, temp)
				}
				return
			}
		}
	}

	k := 3
	for {
		search(2, k-1, 2*nums[k])
		if len(nums) >= n {
			break
		}
		nums = append(nums, 2*nums[k])
		k++
	}

	return nums[n-1]
}
