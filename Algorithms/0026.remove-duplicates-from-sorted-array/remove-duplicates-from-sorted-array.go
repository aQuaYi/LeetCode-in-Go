package problem0026

func removeDuplicates(a []int) int {
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}
	return left + 1
}
