package problem0922

func sortArrayByParityII(A []int) []int {
	size := len(A)
	res := make([]int, size)
	odd, even := 0, 0
	for _, a := range A {
		if a%2 == 0 {
			res[even*2] = a
			even++
		} else {
			res[odd*2+1] = a
			odd++
		}
	}

	return res
}
