package problem0914

func hasGroupsSizeX(deck []int) bool {
	size := len(deck)
	if size < 2 {
		return false
	}

	count := make(map[int]int, size)
	for _, v := range deck {
		count[v]++
	}

	g := count[deck[0]]

	for _, c := range count {
		g = gcd(g, c)
	}

	return g > 1
}

// 最大公约数
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
