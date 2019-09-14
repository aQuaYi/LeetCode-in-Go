package problem1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	total, part := 0, 0
	for i, d := range distance {
		total += d
		if start <= i && i < destination {
			part += d
		}
	}

	return min(part, total-part)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
