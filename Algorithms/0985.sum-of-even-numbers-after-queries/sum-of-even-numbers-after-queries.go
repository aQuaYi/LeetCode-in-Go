package problem0985

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0, len(A))

	sumEven := 0
	for _, v := range A {
		if v%2 == 0 {
			sumEven += v
		}
	}

	for _, q := range queries {
		v, i := q[0], q[1]
		old, new := A[i], A[i]+v
		if old%2 == 0 {
			sumEven -= old
		}
		if new%2 == 0 {
			sumEven += new
		}
		res = append(res, sumEven)
		A[i] = new
	}

	return res
}
