package problem0795

func numSubarrayBoundedMax(a []int, l int, r int) int {
	res, heads, tails := 0, 0, 0

	for i := 0; i < len(a); i++ {
		if l <= a[i] && a[i] <= r {
			heads += tails + 1
			tails = 0
			res += heads
		} else if a[i] < l {
			tails++
			res += heads
		} else {
			heads = 0
			tails = 0
		}
	}

	return res
}
