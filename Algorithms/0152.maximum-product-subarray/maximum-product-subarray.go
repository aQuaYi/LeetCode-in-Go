package Problem0152

func maxProduct(a []int) int {
	Len := len(a)

	// maxs 用来存放最大值
	maxs := make([]int, Len)
	maxs[0] = a[0]
	// mins 用来存放最小值
	mins := make([]int, Len)
	mins[0] = a[0]

	for i := 1; i < Len; i++ {
		if a[i] > 0 {
			maxs[i] = a[i] * maxs[i-1]
			// 因为有可能 masx[i-1] <= 0
			// 所以，需要这个步骤更新答案
			if a[i] > a[i]*maxs[i-1] {
				maxs[i] = a[i]
			}
			mins[i] = a[i] * mins[i-1]
		} else if a[i] < 0 {
			maxs[i] = a[i] * mins[i-1]
			mins[i] = a[i] * maxs[i-1]
			// 因为有可能 maxs[i-1] <= 0
			// 所以，需要这个步骤更新答案
			if a[i] < a[i]*maxs[i-1] {
				mins[i] = a[i]
			}
		}
	}

	res := maxs[0]
	for _, m := range maxs {
		if res < m {
			res = m
		}
	}

	return res
}
