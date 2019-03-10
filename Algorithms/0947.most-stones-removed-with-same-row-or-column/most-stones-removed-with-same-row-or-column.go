package problem0947

func removeStones(stones [][]int) int {
	res := 0
	hasRow := [1000]bool{}
	hasCol := [10000]bool{}
	for _, s := range stones {
		i, j := s[0], s[1]
		if hasRow[i] || hasCol[j] {
			res++
			continue
		}
		hasRow[i] = true
		hasCol[j] = true
	}
	return res
}
