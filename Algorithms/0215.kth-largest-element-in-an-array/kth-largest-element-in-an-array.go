package Problem0215

import "math/rand"

func findKthLargest(a []int, k int) int {
	quickSort(a)
	return a[k-1]
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	i := partition(a)
	quickSort(a[:i])
	quickSort(a[i+1:])
}

func partition(a []int) int {
	size := len(a)
	i := rand.Intn(size)

	a[0], a[i] = a[i], a[0]
	t := a[0]

	lo, hi := 1, size-1
	for {
		for lo < size-1 && a[lo] >= t {
			lo++
		}
		for 0 < hi && t >= a[hi] {
			hi--
		}
		if lo >= hi {
			break
		}
		a[lo], a[hi] = a[hi], a[lo]
	}

	a[hi], a[0] = a[0], a[hi]
	return hi
}
