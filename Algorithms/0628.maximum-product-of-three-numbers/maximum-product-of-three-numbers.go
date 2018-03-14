package problem0628

func maximumProduct(nums []int) int {

	max := -1001
	max1 := -1001
	max2 := -1001
	min1 := 1001
	min2 := 1001

	for _, n := range nums {
		switch {
		case n > max:
			max2, max1, max = max1, max, n
		case n > max1:
			max2, max1 = max1, n
		case n > max2:
			max2 = n
		}

		switch {
		case n < min1:
			min2, min1 = min1, n
		case n < min2:
			min2 = n
		}
	}

	return bigger(max1*max2, min1*min2) * max
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
