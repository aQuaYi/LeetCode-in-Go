package problem0492

import (
	"math"
)

func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i > 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}

	return []int{area, 1}
}
