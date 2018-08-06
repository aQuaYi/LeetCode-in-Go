package problem0873

func lenLongestFibSubseq(a []int) (r int) {
	size := len(a)

	for k := 2; k < size; k++ {
		for i, j := 0, k-1; i < j; {
			s := a[i] + a[j]

			if s < a[k] {
				i++
				continue
			} else if s > a[k] {
				j--
				continue
			}

			count := 3
			x0, x1 := a[j], a[k]
			l := k + 1

			for {

				f, p, x2 := l, size, x0+x1

				for p-f >= 4 {
					m := (p + f) / 2

					if x2 < a[m] {
						p = m
					} else if x2 > a[m] {
						f = m + 1
					} else {
						f = m
						break
					}
				}

				for ; f < p; f++ {
					if a[f] == x2 {
						count, x0, x1, l = count+1, x1, x2, f+1
						break
					}
				}
				if f == p {
					break
				}

			}

			r = max(r, count)
			i++
			j--
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
