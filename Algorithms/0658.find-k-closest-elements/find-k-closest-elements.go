package Problem0658

func findClosestElements(a []int, k int, x int) []int {
	lo, hi := 0, len(a)-k
	for lo < hi {
		mid := lo + (hi-lo)/2
		if x-a[mid] > a[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return a[lo : lo+k]
}
