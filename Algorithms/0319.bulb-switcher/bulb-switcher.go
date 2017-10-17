package Problem0319

func bulbSwitch(n int) int {
	bulbs := make([]bool, n)

	for i := 2; i <= n; i++ {
		for k := i - 1; k < n; k += i {
			bulbs[k] = !bulbs[k]
		}
	}

	var res int
	for i := range bulbs {
		if !bulbs[i] {
			res++
		}
	}

	return res
}
