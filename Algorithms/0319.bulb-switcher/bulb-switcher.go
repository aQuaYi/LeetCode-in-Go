package Problem0319

func bulbSwitch(n int) int {
	res := 1
	if n <= 3 {
		return res
	}

	primes := make([]int, 2)
	primes[0], primes[1] = 2, 3

	var count func(int) int
	count = func(x int) int {
		isSquare := true
		base2 := 1
		for _, p := range primes {
			c := 0
			if p > x {
				break
			}
			for x%p == 0 {
				x /= p
				base2 *= 2
				c++
			}
			if isSquare && c%2 != 0 {
				isSquare = false
			}
		}
		if base2 == 1 {
			primes = append(primes, x)
			return 2
		}
		squareTail := 0
		if isSquare {
			squareTail = 1
		}

		res2 := base2/2 + 2 - squareTail
		return res2
	}

	var i int
	for i = 4; i <= n; i++ {
		if count(i)%2 == 1 {
			res++
		}
	}

	return res
}
