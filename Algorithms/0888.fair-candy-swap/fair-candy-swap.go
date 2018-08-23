package problem0888

func fairCandySwap(A []int, B []int) []int {
	diff := (sum(A) - sum(B)) / 2
	aMap := make(map[int]bool, 100000)
	bMap := make(map[int]bool, 100000)

	i := 0
	aSize := len(A)
	bSize := len(B)

	for {
		if i < aSize {
			aMap[A[i]] = true
			if bMap[A[i]-diff] {
				return []int{A[i], A[i] - diff}
			}
		}

		if i < bSize {
			bMap[B[i]] = true
			if aMap[B[i]+diff] {
				return []int{B[i] + diff, B[i]}
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
