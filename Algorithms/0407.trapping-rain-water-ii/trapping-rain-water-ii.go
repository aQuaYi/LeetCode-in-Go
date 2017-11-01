package Problem0407

func trapRainWater(hs [][]int) int {
	if len(hs) < 3 || len(hs[0]) < 3 {
		return 0
	}
	m, n := len(hs), len(hs[0])

	hh := make([][]int, m)
	vh := make([][]int, m)
	for i := range hh {
		hh[i] = make([]int, n)
		vh[i] = make([]int, n)
		copy(hh[i], hs[i])
		copy(vh[i], hs[i])
	}

	var i, j int

	for i = 1; i < m-1; i++ {
		for j = n - 2; 0 <= j; j-- {
			if hs[i][j] < hs[i][j+1] {
				hh[i][j] = hh[i][j+1]
			}
		}
	}

	for j = 1; j < n-1; j++ {
		for i = m - 2; 0 <= i; i-- {
			if hs[i][j] < hs[i+1][j] {
				vh[i][j] = vh[i+1][j]
			}
		}
	}

	vol := 0
	var wall int

	for i = 1; i < m-1; i++ {
		for j = 1; j < n-1; j++ {
			wall = min4(hh[i-1][j], hh[i+1][j], vh[i][j-1], vh[i][j+1])
			vol += max(wall-hs[i][j], 0)
			hh[i][j] = max(wall, hs[i][j])
			vh[i][j] = max(wall, hs[i][j])
		}
	}

	return vol
}

func min4(a, b, c, d int) int {
	return min(a, min(b, min(c, d)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
