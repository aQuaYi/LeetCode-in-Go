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
		k := math.Floor(math.Pow(num, 1/m))
		if int((math.Pow(k, m+1)-1)/(k-1)) == numInt {
			return strconv.Itoa(int(k))
		}
	}

	return strconv.Itoa(numInt - 1)
}
