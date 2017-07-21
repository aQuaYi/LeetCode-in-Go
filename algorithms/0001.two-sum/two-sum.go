package Problem0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
