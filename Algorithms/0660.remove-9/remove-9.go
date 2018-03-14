package problem0660

func newInteger(n int) int {
	res, base := 0, 1
	for n > 0 {
		res += n % 9 * base
		base *= 10
		n /= 9
	}
	return res
}
