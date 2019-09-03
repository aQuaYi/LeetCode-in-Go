package problem1177

// TODO: unaccepted. the biggest test case spent 4500 ms. Limit Time Exceeded.

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(queries)

	cnt := make([]int, 1, n+1)
	c := 0
	for _, l := range s {
		c ^= 1 << uint(l-'a')
		cnt = append(cnt, c)
	}

	res := make([]bool, n)
	hasSeen := make(map[int]bool, n)
	for i, q := range queries {
		lo, hi, k := q[0], q[1], q[2]
		if k >= 13 {
			res[i] = true
			continue
		}
		m := mark(lo, hi, k)
		r, ok := hasSeen[m]
		if ok {
			res[i] = r
			continue
		}
		remains := bits(cnt[hi+1] ^ cnt[lo])
		r = remains/2 <= k
		res[i] = r
		hasSeen[m] = r
	}

	return res
}

func bits(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}

func mark(lo, hi, k int) int {
	return lo<<22 + hi<<4 + k
}
