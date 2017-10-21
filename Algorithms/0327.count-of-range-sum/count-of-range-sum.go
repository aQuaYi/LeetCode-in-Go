package Problem0327

func countRangeSum(a []int, lower int, upper int) int {
	size := len(a)
	var count = 0
	check := func(n int) {
		if lower <= n && n <= upper {
			count++
		}
	}
	var sum int
	for i := 0; i < size; i++ {
		sum = 0
		for j := i; j < size; j++ {
			sum += a[j ]
			check(sum)
		}
	}

	return count
}
