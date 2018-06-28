package problem0849

func maxDistToClosest(seats []int) int {
	size := len(seats)

	left := make([]int, size)
	right := make([]int, size)

	leftPerson := -size
	rightPerson := size * 2

	for i := 0; i < size; i++ {
		if seats[i] == 1 {
			leftPerson = i
		}
		left[i] = i - leftPerson

		j := size - i - 1
		if seats[j] == 1 {
			rightPerson = j
		}
		right[j] = rightPerson - j
	}

	res := 0
	for i := 0; i < size; i++ {
		res = max(res, min(left[i], right[i]))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
