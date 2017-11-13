package Problem0456

func find132pattern(nums []int) bool {
	ak := -1 << 31
	ajs := make([]int, 0, len(nums))

	for i := len(nums) - 1; 0 <= i; i-- {

		if nums[i] < ak {
			return true
		}

		for len(ajs) > 0 &&
			ajs[len(ajs)-1] < nums[i] {
			ak = ajs[len(ajs)-1]
			ajs = ajs[:len(ajs)-1]
		}

		ajs = append(ajs, nums[i])
	}

	return false
}
