package problem0852

func peakIndexInMountainArray(a []int) int {
	size := len(a)
	i := 1
	for ; i+1 < size; i++ {
		if a[i] > a[i+1] {
			break
		}
	}
	return i
}
