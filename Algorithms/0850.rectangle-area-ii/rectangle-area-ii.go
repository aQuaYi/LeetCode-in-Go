package problem0850

const mod = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	sum := 0
	for _, r := range rectangles {
		sum += area(r)
	}
	return sum % mod
}

func area(r []int) int {
	return ((r[2] - r[0]) * (r[3] - r[1])) % mod
}
