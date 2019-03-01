package problem0945

func minIncrementForUnique(A []int) int {
	counts := [40001]int{}
	max := 0

	for _, n := range A {
		counts[n]++
		if max < n {
			max = n
		}
	}

	res := 0

	for n := 0; n < max; n++ {
		if counts[n] <= 1 { // no redundance
			continue
		}
		redundance := counts[n] - 1
		// move all redundance to n+1
		res += redundance
		counts[n+1] += redundance
	}

	redundance := counts[max] - 1
	res += (redundance + 1) * redundance / 2

	return res
}
