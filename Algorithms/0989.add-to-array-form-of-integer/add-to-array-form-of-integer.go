package problem0989

import "strconv"

func addToArrayForm(A []int, K int) []int {
	a := slice2num(A)
	a += K
	return num2Slice(a)
}

func slice2num(A []int) int {
	res := 0
	for _, v := range A {
		res = res*10 + v
	}
	return res
}

func num2Slice(n int) []int {
	s := strconv.Itoa(n)
	bytes := []byte(s)
	res := make([]int, len(bytes))
	for i, b := range bytes {
		res[i] = int(b - '0')
	}
	return res
}
