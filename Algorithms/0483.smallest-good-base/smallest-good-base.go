package Problem0483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	numInt, _ := strconv.Atoi(n)
	num := float64(numInt)

	mMax := math.Floor(math.Log2(num))

	for m := mMax; 1 < m; m-- {
		k := int(math.Pow(num, 1.0/m))

		tmp := int(math.Pow(float64(k), m+1)-1) / (k - 1)

		if tmp == numInt {
			return strconv.Itoa(k)
		}
	}

	return strconv.Itoa(numInt - 1)
}
