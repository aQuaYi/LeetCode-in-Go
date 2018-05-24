package problem0001

func twoSum(nums []int, target int) []int {
	// m 负责保存map[整数]整数的序列号
	m := make(map[int]int, len(nums))

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := m[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i，因为j<i
		}

		// 把i和i的值，存入map
		m[b] = i
	}

	return nil
}
