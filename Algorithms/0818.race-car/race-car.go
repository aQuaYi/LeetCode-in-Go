package problem0818

import (
	"math"
)

func racecar(target int) int {
	switch target {
	case 1:
		return 1
	case -1:
		return 2
	default:
		return helper(target, 0)
	}
}

func helper(target, pre int) int {
	switch target {
	case 0:
		return pre
	case 1:
		return pre + 1
	}

	s := step(target)

	n := nextTarget(target, s)

	if n < 0 {
		pre += s + 1
		return helper(-n, pre)
	}
	pre += s + 2
	return helper(n, pre)

}

func step(target int) int {
	return int(math.Log2(float64(target))*2+1) / 2
}

func nextTarget(target, step int) int {
	return target -
		(int(math.Pow(2, float64(step))) - 1)
}
