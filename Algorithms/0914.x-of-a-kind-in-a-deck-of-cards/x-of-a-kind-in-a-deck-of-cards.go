package problem0914

func hasGroupsSizeX(deck []int) bool {
	size := len(deck) // 题目说了 size >= 1

	// 统计 deck 中相同数字的数量
	count := make(map[int]int, size)
	for _, card := range deck {
		count[card]++
	}

	d := count[deck[0]]

	for _, c := range count {
		d = gcd(d, c)
		if d == 1 {
			return false
		}
	}

	return true
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
