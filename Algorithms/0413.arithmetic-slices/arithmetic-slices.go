package Problem0413

func numberOfArithmeticSlices(a []int) int {
	if len(a) < 3 {
		return 0
	}

	if a[1]-a[0] != a[2]-a[1] {
		return numberOfArithmeticSlices(a[1:])
	}

	var i = 3
	diff := a[1] - a[0]
	for i < len(a) && a[i]-a[i-1] == diff {
		i++
	}

	return (i*i-3*i+2)>>1 + numberOfArithmeticSlices(a[i:])
}
