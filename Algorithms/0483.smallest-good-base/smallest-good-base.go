package Problem0483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)

	mMax := int(math.Log2(float64(num)))

	for m := mMax; 1 < m; m-- {
		k := int(math.Pow(float64(num), 1.0/float64(m)))

		tmp := int(math.Pow(float64(k), float64(m)+1)-1) / (k - 1)

		if tmp == num {
			return strconv.Itoa(k)
		}
	}

	return strconv.Itoa(num - 1)
}
