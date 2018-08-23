package problem0888

func fairCandySwap(A []int, B []int) []int {
	halfDiff := (sum(A) - sum(B)) / 2
	aIsExist := make(map[int]bool, 1000)
	bIsExist := make(map[int]bool, 1000)

	i := 0
	aSize := len(A)
	bSize := len(B)

	for {
		if i < aSize {
			aIsExist[A[i]] = true
			if bIsExist[A[i]-halfDiff] {
				return []int{A[i], A[i] - halfDiff}
			}
		}

		if i < bSize {
			bIsExist[B[i]] = true
			if aIsExist[B[i]+halfDiff] {
				return []int{B[i] + halfDiff, B[i]}
			}
		}

		i++
	}
}

func sum(a []int) int {
	res := 0
	for _, n := range a {
		res += n
	}
	return res
}
