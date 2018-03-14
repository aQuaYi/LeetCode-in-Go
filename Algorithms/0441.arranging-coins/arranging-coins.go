package problem0441

func arrangeCoins(n int) int {
	n *= 2
	x := n
	for !(x*(x+1) <= n && n < (x+1)*(x+2)) {
		x = (x + n/(x+1)) / 2
	}
	return x
}
