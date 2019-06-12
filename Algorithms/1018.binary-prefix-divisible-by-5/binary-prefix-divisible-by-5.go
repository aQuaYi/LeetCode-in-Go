package problem1018

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	r := 0
	for i, a := range A {
		r = (r*2 + a) % 5
		if r == 0 {
			res[i] = true
		}
	}
	return res
}
