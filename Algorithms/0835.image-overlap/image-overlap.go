package problem0835

func largestOverlap(a [][]int, b [][]int) int {
	res := 0
	size := len(a)

	for m := 0; m < size; m++ {
		for n := 0; n < size; n++ {
			ta := 0
			tb := 0
			for i := 0; i < size && i+m < size; i++ {
				for j := 0; j < size && j+n < size; j++ {
					if a[i][j] == 1 && b[i+m][j+n] == 1 {
						ta++
					}
					if b[i][j] == 1 && a[i+m][j+n] == 1 {
						tb++
					}
				}
			}
			res = max(res, max(ta, tb))
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
