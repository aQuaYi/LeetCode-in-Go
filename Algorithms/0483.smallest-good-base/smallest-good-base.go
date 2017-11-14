package Problem0483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	value, _ := strconv.ParseUint(n, 10, 64)
	mMax := int(math.Log2(float64(value)))

	for m := mMax; m >= 1; m-- {
		cur := search(value, m)
		if cur != 0 {
			return strconv.FormatUint(cur, 10)
		}
	}

	return strconv.FormatUint(value-1, 10)
}

func search(value uint64, m int) uint64 {
	left, right := uint64(1), uint64(math.Pow(float64(value), 1.0/float64(m)))+1
	for left <= right {
		mid := (left + right) / 2
		var sum, cur uint64

		sum, cur = 1, 1
		for i := 1; i <= m; i++ {
			cur *= mid
			sum += cur
		}

		if sum == value {
			return mid
		}
		if sum > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return 0
}
