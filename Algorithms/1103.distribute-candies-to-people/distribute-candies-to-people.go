package problem1103

import "math"

func distributeCandies(candies, people int) []int {
	res := make([]int, people)

	k := root(candies)

	// res[i] = (i+1) + (i+1+p) + (i+1+p*2) + ... + (i+1+n*p)
	// n = (k-i-1)/p , p = people
	candiesOf := func(i int) int {
		c0 := i + 1
		n := (k - c0) / people
		cn := c0 + n*people
		return (n + 1) * (c0 + cn) / 2
	}

	// i<k for in case k<people
	for i := 0; i < people && i < k; i++ {
		res[i] = candiesOf(i)
	}

	// remaining
	res[k%people] += candies - k*(k+1)/2

	return res
}

// root returns k for k*(k+1)/2 = candies
func root(candies int) int {
	delta := float64(1 + 8*candies)
	return int((math.Sqrt(delta) - 1) / 2)
}
