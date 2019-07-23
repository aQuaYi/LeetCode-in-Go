package problem1104

import "math"

func pathInZigZagTree(label int) []int {
	index := math.Floor(math.Log2(float64(label)))
	t := int(math.Pow(2, index))
	// sum is max+min in the above level
	i, sum := int(index), t+t/2-1
	res := make([]int, i+1)
	for i >= 0 {
		res[i] = label
		label = sum - label/2
		sum /= 2
		i--
	}
	return res
}
