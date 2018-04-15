package problem0799

func champagneTower(poured int, queryRow int, queryGlass int) float64 {
	// 答案是对称的，统一求解左边半边
	if queryGlass > queryRow/2 {
		queryGlass = queryRow - queryGlass
	}

	row := [100]float64{0: float64(poured)}
	i := 0 // 代表 row 所在的行数
	for {
		nextRow := [100]float64{}
		for j := max(i+queryGlass-queryRow, 0); j <= queryGlass; j++ {
			if row[j] > 1 {
				// 对于所有超过的杯子，多余的香槟，分给它下面的杯子
				nextRow[j] += (row[j] - 1) / 2
				nextRow[j+1] += (row[j] - 1) / 2
				row[j] = 1
			}
		}

		if i == queryRow {
			// 已经完成了对最后一行的检查
			return row[queryGlass]
		}

		row = nextRow
		i++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
