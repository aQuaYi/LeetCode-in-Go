package problem1250

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func isGoodArray(nums []int) bool {
	val := nums[0]
	for _, num := range nums[1:] {
		val = gcd(val, num)
	}
	return val == 1
}
