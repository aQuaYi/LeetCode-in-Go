package problem0593

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

// 获取以 ps 构成的四边形的中心为起点，到 ps 中各个点的向量
func getVectors(ps [][]int) [][]int {
	vs := make([][]int, 4)
	c := getCenter(ps)
	for i := range vs {
		vs[i] = vector(ps[i], c)
	}
	return vs
}

// NOTICE: 为了避免中心坐标出现小数，将四边形变成原来的 4 倍进行处理
func getCenter(ps [][]int) []int {
	x, y := 0, 0
	for i := 0; i < len(ps); i++ {
		ps[i][0] *= 4
		ps[i][1] *= 4
		x += ps[i][0]
		y += ps[i][1]
	}
	return []int{x / 4, y / 4}
}

func vector(p, c []int) []int {
	return []int{p[0] - c[0], p[1] - c[1]}
}

func isOK(v1, v2 []int) bool {
	if isZero(v1) || isZero(v2) {
		return false
	}
	return isInverse(v1, v2) ||
		(isVertical(v1, v2) && isLengthEqual(v1, v2))
}

func isZero(v []int) bool {
	return v[0] == 0 && v[1] == 0
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
