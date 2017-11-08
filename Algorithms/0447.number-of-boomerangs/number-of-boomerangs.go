package Problem0447

func numberOfBoomerangs(points [][]int) int {
	res := 0

	size := len(points)
	for i := 0; i < size-2; i++ {
		for j := i + 1; j < size-1; j++ {
			for k := j + 1; k < size; k++ {
				pi := points[i]
				pj := points[j]
				pk := points[k]
				if isBoomerang(pi, pj, pk) ||
					isBoomerang(pj, pi, pk) ||
					isBoomerang(pk, pj, pi) {
					res += 2
				}
			}
		}
	}

	return res
}

func isBoomerang(i, j, k []int) bool {
	return dSquare(i, j) == dSquare(i, k)
}

func dSquare(p1, p2 []int) int {
	x := p2[0] - p1[0]
	y := p2[1] - p1[1]
	return x*x + y*y
}
