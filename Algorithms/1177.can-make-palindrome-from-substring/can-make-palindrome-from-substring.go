package problem1177

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(queries)

	cnt := make([]int, 1, n+1)
	c := 0
	for _, l := range s {
		c ^= 1 << uint(l-'a')
		cnt = append(cnt, c)
	}

	res := make([]bool, n)

	for i, q := range queries {
		lo, hi, k := q[0], q[1], q[2]
		remains := bites(cnt[hi+1] ^ cnt[lo])
		res[i] = remains/2 <= k
	}

	return res
}

func bites(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
