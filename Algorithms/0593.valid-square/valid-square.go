package Problem0593

func validSquare(p1, p2, p3, p4 []int) bool {
	vs := getVectors([][]int{p1, p2, p3, p4})
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if !isOK(vs[i], vs[j]) {
				return false
			}
		}
	}
	return true
}

func getVectors(ps [][]int) [][]int {
	vs := make([][]int, 4)
	c := getCenter(ps)
	for i := range vs {
		vs[i] = vector(ps[i], c)
	}
	return vs
}

func getCenter(ps [][]int) []int {
	x, y := 0, 0
	for i := 0; i < len(ps); i++ {
		ps[i][0] *= 2
		ps[i][1] *= 2
		x += ps[i][0]
		y += ps[i][1]
	}
	return []int{x / len(ps), y / len(ps)}
}

func vector(p, c []int) []int {
	return []int{p[0] - c[0], p[1] - c[1]}
}

//

func isOK(v1, v2 []int) bool {
	return isInverse(v1, v2) || (isVertical(v1, v2) && isLengthEqual(v1, v2))
}

func isInverse(v1, v2 []int) bool {
	return v1[0]+v2[0] == 0 && v1[1]+v2[1] == 0
}

func isVertical(v1, v2 []int) bool {
	return v1[0]*v2[0]+v1[1]*v2[1] == 0
}

func isLengthEqual(v1, v2 []int) bool {
	return lenSquare(v1) == lenSquare(v2)
}

func lenSquare(v []int) int {
	return v[0]*v[0] + v[1]*v[1]
}
