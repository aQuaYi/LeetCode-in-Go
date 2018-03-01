package Problem0764

func orderOfLargestPlusSign(N int, mines [][]int) int {
	k := N/2 + N%2

	if len(mines) == 0 {
		return k
	}

	isMined := [500][500]bool{}
	for _, m := range mines {
		isMined[m[0]][m[1]] = true
	}

	isGoodEdge := func(xBegin, xEnd, yBegin, yEnd int) bool {
		for i := xBegin; i < xEnd; i++ {
			for j := yBegin; j < yEnd; j++ {
				if isMined[i][j] {
					return false
				}
			}
		}
		return true
	}

	for ; k > 0; k-- {
		for i := k - 1; i <= N-k; i++ {
			for j := k - 1; j <= N-k; j++ {
				if !isMined[i][j] &&
					isGoodEdge(i-k+1, i, j, j) &&
					isGoodEdge(i+1, i+k, j, j) &&
					isGoodEdge(i, i, j-k+1, j) &&
					isGoodEdge(i, i, j+1, j+k) {
					return k
				}
			}
		}
	}

	return k
}
