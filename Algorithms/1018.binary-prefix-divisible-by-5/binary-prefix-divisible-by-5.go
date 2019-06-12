package problem1018

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	r := 0
	for i, a := range A {
		r = (r*2 + a) % 5
		res[i] = r == 0
	}
	return res
}
