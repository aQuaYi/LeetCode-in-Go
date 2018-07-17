package problem0457

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	if size < 2 {
		return false
	}

	for i := 0; i < size; i++ {
		nums[i] = nums[i] % size
	}

	for i, n := range nums {
		if n == -size || n == 0 || n == size {
			continue
		}

		mark := size
		if n < 0 {
			mark = -size
		}

		for {
			ni := (size + nums[i] + i) % size
			nums[i] = mark
			if nums[ni] == mark {
				return true
			}
			if nums[ni]*mark < 0 {
				break
			}
			i = ni
		}
	}
	return false
}
