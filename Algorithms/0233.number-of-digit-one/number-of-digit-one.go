package Problem0233

func countDigitOne(n int) int {
	ones := 0
	m := 1
	for m <= n {
		a, b := n/m, n%m
		ones += (a + 8) / 10 * m
		if a%10 == 1 {
			ones += b + 1
		}
		m *= 10
	}
	return ones
}
