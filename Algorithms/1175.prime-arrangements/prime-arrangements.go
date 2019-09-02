package problem1175

import "sort"

const mod = 1e9 + 7

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func numPrimeArrangements(n int) int {
	c := count(n)
	return factorial(c) * factorial(n-c) % mod
}

func count(n int) int {
	return sort.Search(25, func(i int) bool {
		return primes[i] > n
	})
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1) % mod
}
