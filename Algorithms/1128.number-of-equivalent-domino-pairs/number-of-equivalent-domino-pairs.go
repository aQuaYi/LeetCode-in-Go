package problem1128

func numEquivDominoPairs(dominoes [][]int) int {
	res := 0
	count := [100]int{}
	for _, domino := range dominoes {
		d := mapping(domino)
		res += count[d]
		count[d]++
	}
	return res
}

func mapping(A []int) int {
	a, b := A[0], A[1]
	if a < b {
		return a*10 + b
	}
	return b*10 + a
}
