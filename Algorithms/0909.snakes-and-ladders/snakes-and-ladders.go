package problem0909

func snakesAndLadders(p [][]int) int {

	return 0
}

func location(m, n, square int) (x, y int) {
	x = m - (square+n-1)/n
	if m%2 == 1 {
		y = (square - 1) % n
	} else {
		y = n - 1 - (square-1)%n
	}
	return
}
