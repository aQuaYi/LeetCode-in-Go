package Problem0441

func arrangeCoins(n int) int {
	x := 2 * n
	for x*(x+1) > 2*n {
		x = (x + 2*n/(x+1)) / 2
	}
	return x
}
