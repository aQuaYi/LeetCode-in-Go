package Problem0376

func wiggleMaxLength(nums []int) int {
	size := len(nums)

	if size < 2 {
		return size
	}

	checkFunc := func(n int) (func(int), func() int) {
		init := 1
		if n < 0 {
			init = -1
		}
		res := 2

		return func(x int) {
				var new int
				switch {
				case x < 0:
					new = -1
				case x > 0:
					new = 1
				default:
					return
				}

				if init*new < 0 {
					res++
					init = new
				}
			}, func() int {
				return res
			}
	}

	var check func(int)
	var res func() int
	var i = 1
	for i < size && nums[i]-nums[i-1] == 0 {
		i++
	}

	if i == size {
		return 1
	}

	check, res = checkFunc(nums[i] - nums[i-1])

	for i < size {
		check(nums[i] - nums[i-1])
		i++
	}

	return res()
}
