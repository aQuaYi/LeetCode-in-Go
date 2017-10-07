package Problem0264

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	hash := make(map[int]bool, n)
	for i := 2; i <= 6; i++ {
		hash[i] = true
	}

	count := 6
	num := 6
	for count < n {
		num++
		if (num%2 == 0 && hash[num/2]) ||
			(num%3 == 0 && hash[num/3]) ||
			(num%5 == 0 && hash[num/5]) {
			hash[num] = true
			count++
		}
	}

	return num
}
