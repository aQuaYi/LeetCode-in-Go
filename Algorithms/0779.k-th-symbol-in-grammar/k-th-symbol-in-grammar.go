package problem0779

func kthGrammar(N int, K int) int {
	return helper(N, K)
}

func helper(n, k int) int {
	row := []int8{0}
	n--
	for n > 0 {
		temp := make([]int8, 0, len(row)*2)
		for i := range row {
			if row[i] == 0 {
				temp = append(temp, 0, 1)
			} else {
				temp = append(temp, 1, 0)
			}
		}
		row = temp
		n--
	}
	return int(row[k-1])
}
