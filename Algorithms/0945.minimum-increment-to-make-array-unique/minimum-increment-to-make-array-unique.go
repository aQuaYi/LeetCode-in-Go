package problem0945

func minIncrementForUnique(A []int) int {
	rec := [80000]bool{}

	rec[A[0]] = true

	size := len(A)
	res := 0

	for i := 1; i < size; i++ {
		n := A[i]
		for rec[n] {
			n++
			res++
		}
		rec[n] = true
	}

	return res
}
