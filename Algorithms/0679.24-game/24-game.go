package Problem0679

func judgePoint24(nums []int) bool {
	fs := ints2float64s(nums)

	left := make([]float64, 2)
	left[0] = fs[0]

	for i := 1; i < 4; i++ {
		left[1] = fs[i]

		right := make([]float64, 0, 2)
		for j := 1; j < 4; j++ {
			if j != i {
				right = append(right, fs[j])
			}
		}

		le := operator(left)
		ri := operator(right)

		// le =?= 24 opt ri
		for r := range ri {
			if le[24+r] ||
				le[24-r] {
				return true
			}

			if r != 0 &&
				(le[24*r] || le[24/r]) {
				return true
			}

		}
	}

	return false
}

func operator(fs []float64) map[float64]bool {
	a, b := fs[0], fs[1]
	res := make(map[float64]bool, 6)

	res[a+b] = true

	res[a-b] = true
	res[b-a] = true

	res[a*b] = true

	res[a/b] = true
	res[b/a] = true

	return res
}

func ints2float64s(nums []int) []float64 {
	res := make([]float64, len(nums))
	for i, n := range nums {
		res[i] = float64(n)
	}
	return res
}
