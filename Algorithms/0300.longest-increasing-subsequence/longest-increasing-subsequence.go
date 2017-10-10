package Problem0300

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))

	for _, n := range nums {
		at := search(tails, n)
		if at == len(tails) {
			tails = append(tails, n)
		} else if tails[at] > n {
			tails[at] = n
		}
	}

	return len(tails)
}

func search(nums []int, n int) int {
	if len(nums) == 0 || n < nums[0] {
		return 0
	}

	if nums[len(nums)-1] < n {
		return len(nums)
	}

	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == n {
			return mid
		} else if n < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
