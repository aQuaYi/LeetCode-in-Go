package problem1042

func gardenNoAdj(N int, paths [][]int) []int {
	connects := make([][]int, N)
	for _, p := range paths {
		i, j := p[0]-1, p[1]-1 // i,j = x-1, y-1
		connects[i] = append(connects[i], j)
		connects[j] = append(connects[j], i)
	}
	res := make([]int, N)
	for i := 0; i < N; i++ {
		isUsed := [5]bool{}
		for _, j := range connects[i] {
			isUsed[res[j]] = true
		}
		for color := 1; color <= 4; color++ {
			if !isUsed[color] {
				res[i] = color
				break
			}
		}
	}
	return res
}
