package problem0045

func jump(nums []int) int {
	i, count, end := 0, 0, len(nums)-1

	var nextI, maxNextI, maxI int
	for i < end {
		if i+nums[i] >= end {
			return count + 1
		}

		nextI, maxNextI = i+1, i+nums[i]
		for nextI <= maxNextI {
			if nextI+nums[nextI] > maxI {
				maxI, i = nextI+nums[nextI], nextI
			}

			nextI++
		}

		count++
	}

	return count
}
