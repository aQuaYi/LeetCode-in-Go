package problem1177

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(queries)

	count := make([][26]int, 1, n+1)
	c := [26]int{}
	for _, letter := range s {
		b := letter - 'a'
		c[b]++
		count = append(count, c)
	}

	res := make([]bool, n)

	for i, q := range queries {
		lo, hi, k := q[0], q[1], q[2]
		single := 0
		for j := 0; j < 26; j++ {
			single += (count[hi+1][j] - count[lo][j]) % 2
			// if single/2 > k {
			// 	break
			// }
		}
		res[i] = single/2 <= k
	}

	return res
}
