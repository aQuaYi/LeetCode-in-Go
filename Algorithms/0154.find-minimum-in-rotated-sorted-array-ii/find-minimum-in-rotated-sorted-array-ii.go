package Problem0154

func findMin(a []int) int {
	Len := len(a)

	i := 1
	for i < Len && a[i-1] <= a[i] {
		i++
	}

	return a[i%Len]
}
