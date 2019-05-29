package problem0986

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0, len(A)+len(B))
	var A0, B0, intersection []int
	for (A0 != nil || len(A) > 0) &&
		(B0 != nil || len(B) > 0) {
		if A0 == nil {
			A0, A = A[0], A[1:]
		}
		if B0 == nil {
			B0, B = B[0], B[1:]
		}
		A0, B0, intersection = intersect(A0, B0)
		if intersection != nil {
			res = append(res, intersection)
		}
	}

	return res
}

func intersect(A0, B0 []int) (a, b, intersection []int) {
	a, b = A0, B0
	if a[0] > b[0] {
		a, b = b, a
		defer func() {
			a, b = b, a
		}()
	}

	if a[1] < b[0] {
		a = nil
		return a, b, nil
	}

	if a[1] == b[1] {
		intersection = []int{b[0], b[1]}
		a, b = nil, nil
		return
	}

	if a[1] < b[1] {
		intersection = []int{b[0], a[1]}
		a, b = nil, []int{a[1], b[1]}
		return
	}

	// a[1]>b[1]
	intersection = []int{b[0], b[1]}
	a, b = []int{b[1], a[1]}, nil
	return
}
