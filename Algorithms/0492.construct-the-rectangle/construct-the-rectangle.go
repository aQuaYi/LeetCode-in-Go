package Problem0492

import (
	"math"
)

func constructRectangle(area int) []int {
	L := int(math.Sqrt(float64(area)))
	W := area / L

	for L*W != area {
		L++
		W = area / L
	}

	return []int{L, W}
}
