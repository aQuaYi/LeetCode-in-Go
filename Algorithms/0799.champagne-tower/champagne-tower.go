package problem0799

func champagneTower(poured int, queryRow int, queryGlass int) float64 {
	// 答案是对称的，统一求解左边半边
	if queryGlass > queryRow/2 {
		queryGlass = queryRow - queryGlass
	}

	oldRow := [99]float64{}
	oldRow[0] = float64(poured)
	for {
		newRow := [99]float64{}
		for j := 0; j <= queryGlass; j++ {
			if oldRow[j] > 1 {
				newRow[j] += (oldRow[j] - 1) / 2
				newRow[j+1] += (oldRow[j] - 1) / 2
				oldRow[j] = 1
			}
		}

		queryRow--
		if queryRow < 0 {
			return oldRow[queryGlass]
		}
		oldRow = newRow
	}

}
