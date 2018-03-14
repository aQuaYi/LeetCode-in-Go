package problem0167

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[target-n] != 0 {
			return []int{m[target-n], i + 1}
		}
		m[n] = i + 1
	}

	return nil
}
