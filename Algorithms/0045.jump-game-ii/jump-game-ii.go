package Problem0045

func jump(nums []int) int {
	// 虽然题目已经假设了 len(nums) > 0
	// 但是，我还是做了长度检查，因为后面使用了切片的索引号，检查一下，养成好习惯。
	if len(nums) == 0 {
		panic("输入nums的长度为0")
	}

	return jumpToNext(nums[:len(nums)-1], 0)
}

func jumpToNext(nums []int, res int) int {
	if len(nums) == 0 {
		return res
	}

	end := 0

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= len(nums) {
			end = i
		}
	}

	return jumpToNext(nums[:end], res+1)
}
