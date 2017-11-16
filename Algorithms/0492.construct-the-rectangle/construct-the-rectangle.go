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

	if L < W {
		L, W = W, L
	}

	return []int{L, W}
}
