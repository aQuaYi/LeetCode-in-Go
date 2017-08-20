package Problem0414

import (
	"math"
)

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		switch {
		case max1 < n:
			max3, max2, max1 = max2, max1, n
		case max2 < n && n != max1:
			max3, max2 = max2, n
		case max3 < n && n != max2 && n != max1:
			max3 = n
		}
	}

	if max3 == math.MinInt64 || max3 == max2 {
		return max1
	}

	return max3
}
