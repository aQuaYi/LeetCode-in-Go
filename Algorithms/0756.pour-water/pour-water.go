package Problem0756

func pourWater(heights []int, V int, K int) []int {
	for V > 0 {
		V--
		drop(heights, K, heights[K]+1, true, true)
	}
	return heights
}

func drop(h []int, i, j int, l, r bool) bool {
	if l && i > 0 && h[i-1] <= h[i] && drop(h, i-1, h[i], l, false) {
		return true
	}
	if r && i < len(h)-1 && h[i+1] <= h[i] && drop(h, i+1, h[i], false, r) {
		return true
	}

	if h[i] == j {
		return false
	}

	h[i]++
	return true
}
