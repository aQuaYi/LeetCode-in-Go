package Problem0440

func findKthNumber(n int, k int) int {
	var res int
	for head := 1; head <= 9; head++ {
		for digits := 1; digits <= 10; digits++ {
			min, max, count := get(head, digits)
			if n < min {
				break
			}
			if n <= max {
				if k <= n-min+1 {
					return min + k - 1
				}
				k -= n - min + 1
				break
			}
			if k <= count {
				return min + k - 1
			}
			k -= count
		}
	}
	return res
}
func get(head, digits int) (min, max, count int) {
	base := 1
	for i := 1; i < digits; i++ {
		base *= 10
	}

	min = head * base
	max = min + base - 1
	count = base
	return
}
