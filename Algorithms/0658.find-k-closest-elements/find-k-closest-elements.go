package problem0658

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	i := sort.Search(len(arr)-k, func(i int) bool { return x-arr[i] <= arr[i+k]-x })
	return arr[i : i+k]
}
