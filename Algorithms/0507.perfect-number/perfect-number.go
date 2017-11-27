package Problem0507

func checkPerfectNumber(num int) bool {
	nextPrime := primeFunc()
	l, r, remain := 1, num, num-1
	p := nextPrime()

	for l < r {
		if num%p != 0 {
			p = nextPrime()
			continue
		}

		for num%p == 0 {
			num /= p
			l *= p
			r = num
			if l >= r {
				break
			}

			remain -= l + r

			if remain < 0 {
				return false
			}
		}
	}

	return remain == 0
}

func primeFunc() func() int {
	r := make([]int, 2, 1024)
	r[0] = 2
	r[1] = 3
	i := -1
	return func() int {
		i++
		if i < len(r) {
			return r[i]
		}
		t := r[i-1] + 2
		j := 0
		for {
			for j = 0; j < i; j++ {
				if t%r[j] == 0 {
					break
				}
			}

			if j == i {
				r = append(r, t)
				return t
			}

			t += 2
		}
	}
}
