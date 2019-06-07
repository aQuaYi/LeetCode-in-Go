package problem1007

func minDominoRotations(A []int, B []int) int {
	a, b := A[0], B[0]
	ca, cb := 1, 1
	for i := 1; i < len(A); i++ {
		if A[i] != a && A[i] != b &&
			B[i] != a && B[i] != b {
			return -1
		}
		if A[i] == a || B[i] == a {
			ca++
		}
		if A[i] == b || B[i] == b {
			cb++
		}
	}

	if ca != len(A) && cb != len(A) {
		return -1
	}

	flag := a
	if ca < cb {
		flag = b
	}

	ca, cb = 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] == flag && B[i] == flag {
			continue
		}
		if A[i] == flag {
			cb++
		} else {
			ca++
		}
	}

	return min(ca, cb)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
