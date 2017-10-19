package Problem0334

func increasingTriplet(nums []int) bool {
	i, j, k := -1<<31, -1<<31, -1<<31
	for _, n := range nums {
		if i < n {
			k, j, i = j, i, n
		} else if j < n {
			i = n
		} else if k < n {
			k, j, i = -1<<31, -1<<31, n
		}

		if k > -1<<31 {
			return true
		}
	}

	return false
}
