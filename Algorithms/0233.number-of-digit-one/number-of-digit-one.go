package problem0233

func countDigitOne(n int) int {
	var a, b, ones int
	m := 1
	for m <= n {
		a, b = n/m, n%m
		ones += (a + 8) / 10 * m
		if a%10 == 1 {
			ones += b + 1
		}
		m *= 10
	}
	return ones
}
