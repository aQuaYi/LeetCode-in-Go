package problem1103

import "math"

func distributeCandies(candies int, num int) []int {
	res := make([]int, num)
	l := length(candies)
	for i := 0; i < num && i < l; i++ {
		res[i] = candy(i, num, l)
	}
	r := candies - l*(l+1)/2
	res[l%num] += r
	return res
}

func length(candies int) int {
	delta := float64(1 + 8*candies)
	l := int((math.Sqrt(delta) - 1) / 2)
	return l
}

func candy(i, d, l int) int {
	a1 := i + 1
	n := (l-a1)/d + 1
	an := a1 + (n-1)*d
	return n * (a1 + an) / 2
}
